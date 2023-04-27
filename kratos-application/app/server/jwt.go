package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	jwtKratos "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv4 "github.com/golang-jwt/jwt/v4"

	"kratos-application/api/gen/v1api"
	"kratos-application/core/jwtx"
)

// JwtMiddleware jwt 验证
func JwtMiddleware(secretKey []byte) middleware.Middleware {
	return selector.Server(
		jwtKratos.Server(func(token *jwtv4.Token) (interface{}, error) {
			return secretKey, nil
		}, jwtKratos.WithClaims(func() jwtv4.Claims {
			return &jwtx.CustomClaims{}
		})),
	).Match(NewWhiteListMatcher()).Build()
}

// NewWhiteListMatcher 验证JWT路由白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[v1api.V1_AuthRegister_FullMethodName] = struct{}{}
	whiteList[v1api.V1_AuthLogin_FullMethodName] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
