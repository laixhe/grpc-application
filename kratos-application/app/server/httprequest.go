package server

import (
	stdHttp "net/http"

	"github.com/gorilla/handlers"
)

// CorsFilter 跨域 cors
func CorsFilter() func(stdHttp.Handler) stdHttp.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		//handlers.ExposedHeaders([]string{"Origin", "Content-Type", "Content-Length", "Authorization"}),
		//handlers.AllowCredentials(),
		//handlers.MaxAge(600),
	)
}
