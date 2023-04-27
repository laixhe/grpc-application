package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
	"go-grpc/comm"
)

// 初始化 Grpc 客户端
func initGrpc() {

	// 连接 GRPC 服务端，并设置拦截器
	conn, err := grpc.Dial(comm.GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(comm.UnaryClientInterceptor))
	if err != nil {
		log.Fatalln(err)
	}
	// 关闭连接
	//defer conn.Close()

	// 初始化 v1 客户端，利用它调用远程方法
	comm.V1Client = v1api.NewV1Client(conn)

	log.Println("初始化 Grpc 客户端成功")
	//select {}
}

// 启动 http 服务
func main() {

	// 初始化 Grpc 客户端
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
