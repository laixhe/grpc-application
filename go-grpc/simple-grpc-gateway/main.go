package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	// 引入 proto 编译生成的包
	"github.com/laixhe/grpc-application/go-grpc/goproto"
)

const (
	// Address gRPC 服务地址
	Address = "0.0.0.0:5051"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := goproto.RegisterUserHandlerFromEndpoint(ctx, mux, Address, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":9090", mux)
}

func main() {
	if err := run(); err != nil {
		fmt.Print(err.Error())
	}
}
