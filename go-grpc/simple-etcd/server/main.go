package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/grpc-application/go-grpc/goproto"

	"github.com/laixhe/grpc-application/go-grpc/simple-etcd/etcd"
)

const (
	// ServeAddress 服务监听地址
	ServeAddress string = ":5051"
	// ServeRegisterAddress 注册服务监听地址
	ServeRegisterAddress string = "localhost:5051"
	// ServeName 服务名称
	ServeName string = "simple_etcd"
)

func main() {

	// 监听地址和端口
	listen, err := net.Listen("tcp", ServeAddress)
	if err != nil {
		fmt.Printf("监听端口失败: %v\n", err)
		return
	}

	// 实例化 grpc Server
	serverGrpc := grpc.NewServer()

	// 注册 User service
	pb.RegisterUserServer(serverGrpc, &User{})

	fmt.Println("开始监听 Grpc 端口 0.0.0.0:5051")

	// 服务注册
	ser, err := etcd.NewServiceRegister([]string{
		"localhost:2379",
	}, ServeName, ServeRegisterAddress, 5)
	if err != nil {
		fmt.Printf("服务注册失败: %v\n", err)
		return
	}
	defer ser.Close()

	// 启动服务
	err = serverGrpc.Serve(listen)
	if err != nil {
		fmt.Println("启动 Grpc 服务失败")
	}
}
