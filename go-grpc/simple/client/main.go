package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"

	// 引入 proto 编译生成的包
	"github.com/laixhe/grpc-application/go-grpc/goproto"
)

const (
	// Address gRPC 服务地址
	Address = "127.0.0.1:5051"
)

var UClient goproto.UserClient

// 一元拦截器
// method 请求方法字符串，req 包含请求的所有信息参数等
// reply 在实际 RPC 调用后存储响应信息，通过 invoker 实际调用
func unaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置处理阶段
	log.Println("method: " + method)
	// 实际的RPC调用
	err := invoker(ctx, method, req, reply, cc, opts...)
	// 后置处理
	log.Println(reply)
	return err
}

// 初始化 Grpc 客户端
func initGrpc() {

	// 连接 GRPC 服务端，并设置拦截器
	conn, err := grpc.Dial(Address, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(unaryClientInterceptor))
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	//defer conn.Close()

	// 初始化 User 客户端，利用它调用远程方法
	UClient = goproto.NewUserClient(conn)

	log.Println("初始化 Grpc 客户端成功")
	//select {}
}

// 启动 http 服务
func main() {

	// 初始化 Grpc 客户端
	initGrpc()

	http.HandleFunc("/user/get", GetUser)
	http.HandleFunc("/user/list", GetUserList)

	log.Println("开始监听 http 端口 0.0.0.0:5050")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Printf("http.ListenAndServe err:%v", err)
	}
}
