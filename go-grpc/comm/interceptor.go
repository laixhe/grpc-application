package comm

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// UnaryClientInterceptor 一元拦截器-客户端
// method 请求方法字符串，req 包含请求的所有信息参数等
// reply 在实际 RPC 调用后存储响应信息，通过 invoker 实际调用
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 前置处理阶段
	log.Println("[unary interceptor method]:", method)
	// 实际的RPC调用
	err := invoker(ctx, method, req, reply, cc, opts...)
	// 后置处理
	log.Println("[unary interceptor reply]:", reply)
	return err
}

// UnaryServerInterceptor 一元拦截器-服务端
// 函数名无特殊要求，参数需一致
// req 包含请求的所有信息，info 包含一元RPC服务的所有信息
func UnaryServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 前置处理逻辑
	log.Printf("[unary interceptor request] %s", info.FullMethod)
	// 完成RPC服务的正常执行
	m, err := handler(ctx, req)
	// 后置处理逻辑
	log.Printf("[unary interceptor resonse] %s", m)
	// 返回响应
	return m, err
}
