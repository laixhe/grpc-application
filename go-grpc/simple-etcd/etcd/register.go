package etcd

import (
	"context"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ServiceRegister 创建租约注册服务
type ServiceRegister struct {
	cli           *clientv3.Client
	leaseID       clientv3.LeaseID                        // 租约ID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse // 租约 KeepAlive 相应 chan
	key           string
	val           string
}

// NewServiceRegister 新建注册服务
func NewServiceRegister(endpoints []string, serviceName, serverAddr string, leaseTtl int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	ser := &ServiceRegister{
		cli: cli,
		key: "/" + schema + "/" + serviceName + "/" + serverAddr,
		val: serverAddr,
	}

	// 申请租约设置时间 keepalive
	err = ser.putKeyWithLease(leaseTtl)
	if err != nil {
		return nil, err
	}

	return ser, nil
}

// 设置租约
func (s *ServiceRegister) putKeyWithLease(leaseTtl int64) error {
	// 创建租约时间
	leaseResp, err := s.cli.Grant(context.Background(), leaseTtl)
	if err != nil {
		return err
	}
	// 注册服务并绑定租约(将服务地址注册到 etcd 中)
	_, err = s.cli.Put(context.Background(), s.key, s.val, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}
	// 建立长连接，设置续租 定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), leaseResp.ID)

	if err != nil {
		return err
	}
	s.leaseID = leaseResp.ID
	s.keepAliveChan = leaseRespChan
	return nil
}

// revoke 取消租约 ???
func (s *ServiceRegister) revoke(ctx context.Context) error {
	_, err := s.cli.Revoke(ctx, s.leaseID)
	return err
}

// ListenLease 监听续租情况 ???
func (s *ServiceRegister) ListenLease(ctx context.Context) {
	for {
		select {
		case _, ok := <-s.keepAliveChan:
			if !ok {
				err := s.revoke(ctx)
				if err != nil {
					log.Println("ListenLease revoke", err)
				}
				log.Println("ListenLease 关闭续租")
				return
			}
			log.Println("ListenLease 续租情况 ?")
		}
	}
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
	// 撤销租约
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	return s.cli.Close()
}
