package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// 引入 proto 编译生成的包
	"github.com/laixhe/grpc-application/go-grpc/goproto"
)

// 一元拦截器
// 函数名无特殊要求，参数需一致
// req 包含请求的所有信息，info 包含一元RPC服务的所有信息
func unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 前置处理逻辑
	log.Printf("[unary interceptor request] %s", info.FullMethod)
	// 完成RPC服务的正常执行
	m, err := handler(ctx, req)
	// 后置处理逻辑
	log.Printf("[unary interceptor resonse] %s", m)
	// 返回响应
	return m, err
}

const (
	// Address gRPC 服务地址
	Address = "0.0.0.0:5051"
)

func main() {
	log.Println("开始监听 Grpc 端口 " + Address)

	// 监听地址和端口
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("监听端口失败: %v", err)
	}

	// 实例化 grpc 服务器
	// 在创建 gRPC 服务器实例的时候注册拦截器
	// NewServer 可传入多个拦截器
	serverGrpc := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor))

	// 将服务注册到 grpc 服务器上
	// 为 User 服务注册业务实现，将 User 服务绑定到 RPC 服务器上
	goproto.RegisterUserServer(serverGrpc, &User{})

	// 注册反射服务，这个服务是 CLI 使用的，跟服务本身没有关系
	reflection.Register(serverGrpc)

	// 绑定 grpc 服务器到指定 tcp
	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		log.Println("启动 Grpc 服务失败")
	}
}
