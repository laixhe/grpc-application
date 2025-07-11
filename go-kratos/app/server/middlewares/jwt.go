package middlewares

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"go-kratos/api/gen/pbauth"
	"go-kratos/core"
)

type AuthorizationContextKey struct{}
type IsAuthorizationContextKey struct{}

const JwtContextKey = "jwt"

// JwtClaims 可根据业务自行添加字段
type JwtClaims struct {
	Uid int `json:"uid"`
	jwtv5.RegisteredClaims
}

func NewJwtClaims(uid, expireTime int) *JwtClaims {
	custom := &JwtClaims{
		Uid: uid,
	}
	nowTime := time.Now()
	custom.ExpiresAt = jwtv5.NewNumericDate(nowTime.Add(time.Duration(expireTime) * time.Second)) // 过期时间
	custom.IssuedAt = jwtv5.NewNumericDate(nowTime)                                               // 发布时间（创建时间）
	custom.NotBefore = jwtv5.NewNumericDate(nowTime)                                              // 生效时间
	return custom
}

func (claims *JwtClaims) GetUid() int {
	return claims.Uid
}

// JwtMiddleware 设置 JWT
func JwtMiddleware(conf *core.Config) middleware.Middleware {
	return selector.Server(
		jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return []byte(conf.Jwt.SecretKey), nil
		},
			jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &JwtClaims{}
			})),
	).Match(NewWhiteListMatcher()).Build()
}

// NewWhiteListMatcher 验证 JWT 路由白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList[pbauth.SAuth_Register_FullMethodName] = struct{}{}
	whiteList[pbauth.SAuth_Login_FullMethodName] = struct{}{}
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
