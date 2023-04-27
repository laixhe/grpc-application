package service

import (
	"context"

	"kratos-application/api/gen/v1api"
)

// AuthRegister 注册用户
//
// @Summary	注册用户
// @Accept   json
// @Produce  json
// @Param    body   body      v1api.AuthRegisterRequest   ture "请求body参数"
// @Success  200    {object}  v1api.AuthRegisterResponse
// @Router   /v1/auth/register [post]
func (s *V1Service) AuthRegister(ctx context.Context, req *v1api.AuthRegisterRequest) (*v1api.AuthRegisterResponse, error) {
	return s.biz.AuthRegister(ctx, req)
}

// AuthLogin 登录用户
//
// @Summary	登录用户
// @Accept   json
// @Produce  json
// @Param    body   body      v1api.AuthLoginRequest   ture "请求body参数"
// @Success  200    {object}  v1api.AuthLoginResponse
// @Router   /v1/auth/login [post]
func (s *V1Service) AuthLogin(ctx context.Context, req *v1api.AuthLoginRequest) (*v1api.AuthLoginResponse, error) {
	return s.biz.AuthLogin(ctx, req)
}

// AuthRefresh 刷新Jwt
//
// @Summary	刷新Jwt
// @Accept   json
// @Produce  json
// @Param Authorization header string false "Bearer token令牌"
// @Param    body   body      v1api.AuthRefreshRequest   ture "请求body参数"
// @Success  200    {object}  v1api.AuthRefreshResponse
// @Router   /v1/auth/refresh [post]
func (s *V1Service) AuthRefresh(ctx context.Context, req *v1api.AuthRefreshRequest) (*v1api.AuthRefreshResponse, error) {
	return s.biz.AuthRefresh(ctx, req)
}
