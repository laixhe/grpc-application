package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/grpc-application/go-grpc/goproto"

	"github.com/laixhe/grpc-application/go-grpc/simple-etcd/etcd"
)

const (
	// Address gRPC 服务地址
	Address = "127.0.0.1:5051"
	// ServeName 服务名称
	ServeName string = "simple_etcd"
)

var UClient pb.UserClient

// 初始化 Grpc 客户端
func initGrpc() {

	grpcEtcd := etcd.NewServiceDiscovery([]string{
		"localhost:2379",
	})
	resolver.Register(grpcEtcd)

	// 连接 GRPC 服务端
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", grpcEtcd.Scheme(), ServeName),
		grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	//defer conn.Close()

	// 初始化 User 客户端
	UClient = pb.NewUserClient(conn)

	log.Println("初始化 Grpc 客户端成功")
	//select {}
}

// 启动 http 服务
func main() {

	initGrpc()

	http.HandleFunc("/user/get", GetUser)
	http.HandleFunc("/user/list", GetUserList)

	log.Println("开始监听 http 端口 0.0.0.0:5050")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Printf("http.ListenAndServe err:%v", err)
	}
}
