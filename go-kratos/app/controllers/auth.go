package controllers

import (
	"context"

	"go-kratos/api/gen/pbauth"
	"go-kratos/app/services"
	"go-kratos/core"
)

type Auth struct {
	pbauth.UnimplementedSAuthServer
	server  *core.Server
	service *services.Service
}

func newAuth(server *core.Server, service *services.Service) *Auth {
	return &Auth{
		server:  server,
		service: service,
	}
}

// Register 注册
func (c *Auth) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	return &pbauth.RegisterResponse{}, nil
}

// Login 登录
func (c *Auth) Login(ctx context.Context, req *pbauth.LoginRequest) (*pbauth.LoginResponse, error) {
	return &pbauth.LoginResponse{}, nil
}

// AuthRefresh 刷新Jwt
func (c *Auth) AuthRefresh(ctx context.Context, req *pbauth.RefreshRequest) (*pbauth.RefreshResponse, error) {
	return &pbauth.RefreshResponse{}, nil
}
