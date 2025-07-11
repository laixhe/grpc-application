package services

import (
	"context"

	"go-kratos/api/gen/pbauth"
	"go-kratos/app/model/dao"
	"go-kratos/core"
)

type Auth struct {
	server *core.Server
	dao    *dao.Dao
}

func NewAuth(server *core.Server, modelDao *dao.Dao) *Auth {
	return &Auth{
		server: server,
		dao:    modelDao,
	}
}

// Register 注册
func (s *Auth) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	return &pbauth.RegisterResponse{}, nil
}

// Login 登录
func (s *Auth) Login(ctx context.Context, req *pbauth.LoginRequest) (*pbauth.LoginResponse, error) {
	return &pbauth.LoginResponse{}, nil
}

// AuthRefresh 刷新Jwt
func (s *Auth) AuthRefresh(ctx context.Context, req *pbauth.RefreshRequest) (*pbauth.RefreshResponse, error) {
	return &pbauth.RefreshResponse{}, nil
}
