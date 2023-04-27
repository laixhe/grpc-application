package comm

import (
	"context"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
)

// V1 定义并实现约定的接口
type V1 struct {
	v1api.UnimplementedV1Server

	ServeRegisterAddress string
}

func NewV1(serveRegisterAddress string) *V1 {
	return &V1{
		ServeRegisterAddress: serveRegisterAddress,
	}
}

func (s *V1) AuthRegister(cxt context.Context, req *v1api.AuthRegisterRequest) (*v1api.AuthRegisterResponse, error) {
	return &v1api.AuthRegisterResponse{
		Token: "AuthRegister Token" + s.ServeRegisterAddress,
		User: &v1api.User{
			Uid:   10,
			Uname: "AuthRegister " + req.Password,
			Email: "AuthRegister " + req.Email,
		},
	}, nil
}
func (s *V1) AuthLogin(cxt context.Context, req *v1api.AuthLoginRequest) (*v1api.AuthLoginResponse, error) {
	return &v1api.AuthLoginResponse{
		Token: "AuthLogin Token" + s.ServeRegisterAddress,
		User: &v1api.User{
			Uid:   10,
			Uname: "AuthLogin " + req.Password,
			Email: "AuthLogin " + req.Email,
		},
	}, nil
}
func (s *V1) AuthRefresh(cxt context.Context, req *v1api.AuthRefreshRequest) (*v1api.AuthRefreshResponse, error) {
	return &v1api.AuthRefreshResponse{
		Token: "AuthRefresh Token" + s.ServeRegisterAddress,
	}, nil
}
