package service

import (
	"context"

	"kratos-application/api/gen/auth"
	"kratos-application/app/biz"
)

type AuthService struct {
	auth.UnimplementedSAuthServer
	authBiz *biz.AuthBiz
}

func NewAuthService() *AuthService {
	return &AuthService{
		authBiz: biz.NewAuthBiz(),
	}
}

// RegisterLogin 注册用户
//
// @Summary	注册用户
// @Accept   json
// @Produce  json
// @Param    body   body      auth.RegisterLoginRequest   ture "请求body参数"
// @Success  200    {object}  auth.RegisterLoginResponse
// @Router   /auth/register_login [post]
func (s *AuthService) RegisterLogin(ctx context.Context, req *auth.RegisterLoginRequest) (*auth.RegisterLoginResponse, error) {
	return s.authBiz.RegisterLogin(ctx, req)
}

// Refresh 刷新Jwt
//
// @Summary	刷新Jwt
// @Accept   json
// @Produce  json
// @Param Authorization header string false "Bearer token令牌"
// @Param    body   body      auth.RefreshRequest   ture "请求body参数"
// @Success  200    {object}  auth.RefreshResponse
// @Router   /auth/refresh [post]
func (s *AuthService) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	return s.authBiz.AuthRefresh(ctx, req)
}
