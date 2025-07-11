package middlewares

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/laixhe/gonet/jwt"
)

// HandlerMiddleware 设置头部
func HandlerMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 设置是否有鉴权
				authorization := tr.RequestHeader().Get(jwt.Authorization)
				isAuthorization := false
				if len(authorization) > 0 {
					if strings.HasPrefix(authorization, jwt.Bearer) {
						isAuthorization = true
					}
				}
				ctx = context.WithValue(ctx, IsAuthorizationContextKey{}, isAuthorization)
				// 设置IP
				ctx = context.WithValue(ctx, IPContextKey{}, tr.RequestHeader().Get(IPHeaderKey))
				//
			}
			return handler(ctx, req)
		}
	}
}
