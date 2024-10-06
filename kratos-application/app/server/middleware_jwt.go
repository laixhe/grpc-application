package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"kratos-application/api/gen/auth"
	"kratos-application/core/conf"
	"kratos-application/core/jwtx"
)

// JwtMiddleware 设置 JWT
func JwtMiddleware() middleware.Middleware {
	jwtx.Checking(conf.Get().GetJwt())
	return selector.Server(
		jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return []byte(conf.Get().GetJwt().GetSecretKey()), nil
		},
			jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &jwtx.CustomClaims{}
			})),
	).Match(NewWhiteListMatcher()).Build()
}

// NewWhiteListMatcher 验证 JWT 路由白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[auth.SAuth_RegisterLogin_FullMethodName] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		// 设置是否有鉴权
		//if operation == user.UserList_FullMethodName {
		//	isAuthorization, ok := ctx.Value(jwtx.IsAuthorizationContextKey{}).(bool)
		//	if !ok {
		//		return false
		//	}
		//	return isAuthorization
		//}
		//
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
