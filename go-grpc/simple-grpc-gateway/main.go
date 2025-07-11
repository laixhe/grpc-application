package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
	"go-grpc/comm"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	//opts := []grpc.DialOption{
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//}
	//err := v1api.RegisterV1HandlerFromEndpoint(ctx, mux, Address, opts)
	err := v1api.RegisterV1HandlerServer(ctx, mux, comm.NewV1(comm.HttpAddr))
	if err != nil {
		return err
	}

	log.Println("开始监听 http 端口", comm.HttpAddr)
	return http.ListenAndServe(comm.HttpAddr, mux)
}

func main() {
	if err := run(); err != nil {
		fmt.Print(err.Error())
	}
}
