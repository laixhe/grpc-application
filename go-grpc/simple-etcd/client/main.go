package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
	"go-grpc/comm"
	"go-grpc/simple-etcd/etcd"
)

// 初始化 Grpc 客户端
func initGrpc() {

	grpcEtcd := etcd.NewServiceDiscovery([]string{
		comm.EtcdAddr,
	})
	resolver.Register(grpcEtcd)

	// 连接 GRPC 服务端
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", grpcEtcd.Scheme(), comm.ServeName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	//defer conn.Close()

	// 初始化 v1 客户端
	comm.V1Client = v1api.NewV1Client(conn)

	log.Println("初始化 Grpc 客户端成功")
	//select {}
}

// 启动 http 服务
func main() {

	initGrpc()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := `
<!DOCTYPE html>
<body>
	<div> post /auth/register email password </div>
	<div> post /auth/login email password </div>
</body>
</html>
`
		_, _ = w.Write([]byte(str))
	})
	http.HandleFunc("/auth/register", comm.AuthRegister)
	http.HandleFunc("/auth/login", comm.AuthLogin)

	log.Println("开始监听 http 端口", comm.HttpAddr)
	err := http.ListenAndServe(comm.HttpAddr, nil)
	if err != nil {
		log.Printf("http.ListenAndServe err:%v", err)
	}
}
