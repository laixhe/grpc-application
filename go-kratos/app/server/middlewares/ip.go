package middlewares

import (
	"net"
	"net/http"
	"strings"

	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
)

type IPContextKey struct{}

const IPHeaderKey = "ip"

// GetIpFilter 获取 IP
func GetIpFilter() kratosHttp.FilterFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			// r.Header.Add(requestx.IPHeaderKey, ip)
			next.ServeHTTP(w, r)
		})
	}
}
