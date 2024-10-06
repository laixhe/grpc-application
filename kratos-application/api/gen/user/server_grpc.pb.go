// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/user/server.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SUser_UserInfo_FullMethodName = "/user.SUser/UserInfo"
)

// SUserClient is the client API for SUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SUserClient interface {
	// 用户信息
	UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
}

type sUserClient struct {
	cc grpc.ClientConnInterface
}

func NewSUserClient(cc grpc.ClientConnInterface) SUserClient {
	return &sUserClient{cc}
}

func (c *sUserClient) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserInfoResponse)
	err := c.cc.Invoke(ctx, SUser_UserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SUserServer is the server API for SUser service.
// All implementations must embed UnimplementedSUserServer
// for forward compatibility.
type SUserServer interface {
	// 用户信息
	UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error)
	mustEmbedUnimplementedSUserServer()
}

// UnimplementedSUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSUserServer struct{}

func (UnimplementedSUserServer) UserInfo(context.Context, *UserInfoRequest) (*UserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSUserServer) mustEmbedUnimplementedSUserServer() {}
func (UnimplementedSUserServer) testEmbeddedByValue()               {}

// UnsafeSUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SUserServer will
// result in compilation errors.
type UnsafeSUserServer interface {
	mustEmbedUnimplementedSUserServer()
}

func RegisterSUserServer(s grpc.ServiceRegistrar, srv SUserServer) {
	// If the following call pancis, it indicates UnimplementedSUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SUser_ServiceDesc, srv)
}

func _SUser_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SUserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SUser_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SUserServer).UserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SUser_ServiceDesc is the grpc.ServiceDesc for SUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.SUser",
	HandlerType: (*SUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfo",
			Handler:    _SUser_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/server.proto",
}