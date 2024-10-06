package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
	"go-grpc/comm"
)

func main() {
	log.Println("开始监听 Grpc 端口 " + comm.GrpcAddr)

	// 监听地址和端口
	listen, err := net.Listen("tcp", comm.GrpcAddr)
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 实例化 grpc 服务器
	// 在创建 gRPC 服务器实例的时候注册拦截器
	// NewServer 可传入多个拦截器
	serverGrpc := grpc.NewServer(grpc.UnaryInterceptor(comm.UnaryServerInterceptor))

	// 将服务注册到 grpc 服务器上
	// 为 V1 服务注册业务实现，将 V1 服务绑定到 RPC 服务器上
	v1api.RegisterV1Server(serverGrpc, comm.NewV1(comm.GrpcAddr))

	// 绑定 grpc 服务器到指定 tcp
	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		log.Println("启动 Grpc 服务失败")
	}
}
