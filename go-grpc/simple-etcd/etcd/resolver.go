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

// ServiceResolver 服务发现
type ServiceResolver struct {
	cli        *clientv3.Client //etcd client
	clientConn resolver.ClientConn
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

	return &ServiceResolver{
		cli: cli,
	}
}

// Build 构建解析器 grpc.Dial()时调用
func (s *ServiceResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	s.clientConn = clientConn

	prefix := "/" + target.URL.Scheme + "/" + target.Endpoint() + "/"
	log.Println("T--- build prefix", prefix)
	// 根据前缀获取现有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}

	//s.clientConn.NewAddress(addrList)
	err = s.clientConn.UpdateState(resolver.State{Addresses: s.getServices()})
	if err != nil {
		log.Println("Build UpdateState", err)
	}

	// 监听服务地址列表的变化
	go s.watcher(prefix)
	return s, nil
}

// ResolveNow 监视(watch)有变化调用
func (s *ServiceResolver) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow", rn)
}

// Scheme return schema
func (s *ServiceResolver) Scheme() string {
	return schema
}

// Close 解析器关闭时调用
func (s *ServiceResolver) Close() {
	s.cli.Close()
}

// watcher 监听前缀
func (s *ServiceResolver) watcher(keyPrefix string) {
	rch := s.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
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
func (s *ServiceResolver) SetServiceList(key, val string) {
	log.Println("SetServiceList", key, val)
	s.serverList.Store(key, resolver.Address{Addr: val})
	err := s.clientConn.UpdateState(resolver.State{Addresses: s.getServices()})
	if err != nil {
		log.Println("SetServiceList UpdateState", err)
	}
}

// DelServiceList 删除服务地址
func (s *ServiceResolver) DelServiceList(key string) {
	log.Println("DelServiceList", key)
	s.serverList.Delete(key)
	err := s.clientConn.UpdateState(resolver.State{Addresses: s.getServices()})
	if err != nil {
		log.Println("DelServiceList UpdateState", err)
	}
}

// GetServices 获取服务地址
func (s *ServiceResolver) getServices() []resolver.Address {
	addrs := make([]resolver.Address, 0, 10)
	s.serverList.Range(func(k, v interface{}) bool {
		addrs = append(addrs, v.(resolver.Address))
		return true
	})
	return addrs
}
