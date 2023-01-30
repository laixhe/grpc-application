package etcd

import (
	"context"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

const schema = "grpc"

// ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli        *clientv3.Client //etcd client
	cc         resolver.ClientConn
	serverList sync.Map //服务列表
}

// NewServiceDiscovery  新建发现服务
func NewServiceDiscovery(endpoints []string) resolver.Builder {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceDiscovery{
		cli: cli,
	}
}

//Build 为给定目标创建一个新的`resolver`，当调用`grpc.Dial()`时执行
func (s *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	log.Println("Build")
	s.cc = cc
	prefix := "/" + target.Scheme + "/" + target.Endpoint + "/"
	// 根据前缀获取现有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}
	s.cc.UpdateState(resolver.State{Addresses: s.getServices()})
	//监视前缀，修改变更的server
	go s.watcher(prefix)
	return s, nil
}

// ResolveNow 监视目标更新
func (s *ServiceDiscovery) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}

// Scheme return schema
func (s *ServiceDiscovery) Scheme() string {
	return schema
}

// Close 关闭
func (s *ServiceDiscovery) Close() {
	s.cli.Close()
}

// watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: // 新增或修改
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: // 删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

// SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.serverList.Store(key, resolver.Address{Addr: val})
	s.cc.UpdateState(resolver.State{Addresses: s.getServices()})
}

// DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.serverList.Delete(key)
	s.cc.UpdateState(resolver.State{Addresses: s.getServices()})
}

// GetServices 获取服务地址
func (s *ServiceDiscovery) getServices() []resolver.Address {
	addrs := make([]resolver.Address, 0, 10)
	s.serverList.Range(func(k, v interface{}) bool {
		addrs = append(addrs, v.(resolver.Address))
		return true
	})
	return addrs
}
