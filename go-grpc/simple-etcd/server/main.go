package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
	"go-grpc/comm"
	"go-grpc/simple-etcd/etcd"
)

func main() {

	// 监听地址和端口
	listen, err := net.Listen("tcp", comm.GrpcAddr)
	if err != nil {
		fmt.Printf("监听端口失败: %v\n", err)
		return
	}
	defer listen.Close()

	// 实例化 grpc Server
	serverGrpc := grpc.NewServer()
	defer serverGrpc.GracefulStop()

	// 注册到 grpc 服务中 User
	v1api.RegisterV1Server(serverGrpc, comm.NewV1(comm.GrpcAddr))

	fmt.Println("开始监听 Grpc 端口", comm.GrpcAddr)

	// 将服务地址注册到 etcd 中
	etcdServer, err := etcd.NewServiceRegister([]string{
		comm.EtcdAddr,
	}, comm.ServeName, comm.GrpcAddr, 5)
	if err != nil {
		fmt.Printf("服务注册失败: %v\n", err)
		return
	}
	defer etcdServer.Close()

	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		fmt.Println("启动 Grpc 服务失败")
	}
}
