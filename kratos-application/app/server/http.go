package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerFiles "github.com/swaggo/files"
	swaggerGin "github.com/swaggo/gin-swagger"

	"kratos-application/api/gen/v1api"
	"kratos-application/app/service"
	"kratos-application/core/config"
	"kratos-application/core/jwtx"
)

func NewHTTPServer(c config.Server, cjwt *jwtx.Config, logger log.Logger) *http.Server {
	var opts = make([]http.ServerOption, 0, 10)
	// config
	opts = append(opts, http.Address(c.Addr()))
	if c.Timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Timeout)*time.Second))
	}
	// filter
	opts = append(opts, http.Filter(
		CorsFilter(), // 跨域 cors
	))
	// middleware
	opts = append(opts, http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		JwtMiddleware([]byte(cjwt.SecretKey)), // jwt 验证
		ValidateMiddleware(),                  // 参数验证
	))
	// response
	opts = append(opts, http.ResponseEncoder(ResponseEncoder())) // http 正确-返回统一结构
	opts = append(opts, http.ErrorEncoder(ErrorEncoder()))       // http 错误-返回统一结构
	// http 服务
	srv := http.NewServer(opts...)
	// gin
	// gin swagger doc ( 由于 github.com/go-kratos/swagger-api 的 doc 描述太复杂了，这里采用 gin swagger 来做 API 文档)
	r := gin.New()
	r.GET("/swagger/*any", swaggerGin.WrapHandler(swaggerFiles.Handler))
	srv.HandlePrefix("/swagger/", r.Handler()) // 以 xxx 为前缀的路由交给 gin 框架来处理
	// v1 http 服务
	v1api.RegisterV1HTTPServer(srv, service.NewV1Service(logger))
	return srv
}
