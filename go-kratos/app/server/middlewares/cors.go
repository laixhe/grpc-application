package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CorsFilter 跨域 cors
func CorsFilter() func(next http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Accept", "Authorization"}),
		handlers.AllowCredentials(),
		handlers.MaxAge(10080),
	)
}
