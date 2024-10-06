package server

import (
	"context"
	"net"
	stdHttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	"kratos-application/core/i18nx"
	"kratos-application/core/jwtx"
	"kratos-application/core/requestx"
)

// CorsFilter 跨域 cors
func CorsFilter() func(next stdHttp.Handler) stdHttp.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
		handlers.MaxAge(10080),
	)
}

// GetIpFilter 获取 IP
func GetIpFilter() http.FilterFunc {
	return func(next stdHttp.Handler) stdHttp.Handler {
		return stdHttp.HandlerFunc(func(w stdHttp.ResponseWriter, r *stdHttp.Request) {
			ip := ""
			// 但在代理环境下，客户端IP可能被封装在 X-Forwarded-For 头部
			// 需要检查并解析该头部以获取客户端的真实 IP 地址
			forwardedFor := r.Header.Get("X-Forwarded-For")
			if forwardedFor != "" {
				// X-Forwarded-For 可能包含多个 IP 地址，最左边的是离客户端最近的 IP
				ip = strings.Split(forwardedFor, ", ")[0]
			}
			if forwardedFor == "" {
				// 如果 X-Forwarded-For 为空，尝试从 X-Real-IP 头部获取 IP 地址
				ip = r.Header.Get("X-Real-IP")
			}
			if ip == "" {
				// 获取请求头中的 Remote Address 它通常包含客户端IP和端口号
				ip = strings.Split(r.RemoteAddr, ":")[0]
			}
			netIP := net.ParseIP(ip)
			if netIP == nil {
				w.WriteHeader(404)
				return
			}
			r.Header.Add(requestx.IPHeaderKey, ip)
			next.ServeHTTP(w, r)
		})
	}
}

// HandlerMiddleware 设置头部
func HandlerMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// 设置语言
				acceptLanguage := tr.RequestHeader().Get(i18nx.LanguageHeaderKey)
				ctx = context.WithValue(ctx, i18nx.LanguageContextKey{}, acceptLanguage)
				// 设置是否有鉴权
				authorization := tr.RequestHeader().Get(jwtx.Authorization)
				isAuthorization := false
				if len(authorization) > 0 {
					if strings.HasPrefix(authorization, jwtx.Bearer) {
						isAuthorization = true
					}
				}
				ctx = context.WithValue(ctx, jwtx.IsAuthorizationContextKey{}, isAuthorization)
				// 设置IP
				ctx = context.WithValue(ctx, requestx.IPContextKey{}, tr.RequestHeader().Get(requestx.IPHeaderKey))
				//
			}
			return handler(ctx, req)
		}
	}
}
