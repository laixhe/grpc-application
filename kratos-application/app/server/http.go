package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	swaggerFiles "github.com/swaggo/files"
	swaggerGin "github.com/swaggo/gin-swagger"

	"kratos-application/api/gen/auth"
	"kratos-application/api/gen/user"
	"kratos-application/app/service"
	"kratos-application/core/conf"
)

func NewHTTPServer() *http.Server {
	if conf.Get().GetServers() == nil || conf.Get().GetServers().GetHttp() == nil {
		panic(errors.New("config servers http is nil"))
	}
	//
	ip := conf.Get().GetServers().GetHttp().GetIp()
	port := conf.Get().GetServers().GetHttp().GetPort()
	timeout := conf.Get().GetServers().GetHttp().GetTimeout()
	log.Debugf("http server Config=%v", conf.Get().GetServers().GetHttp())
	//
	var opts = make([]http.ServerOption, 0, 10)
	// config
	var addr = fmt.Sprintf("%s:%d", ip, port)
	opts = append(opts, http.Address(addr))
	if timeout.IsValid() {
		if timeout.Seconds > 0 || timeout.Nanos > 0 {
			opts = append(opts, http.Timeout(timeout.AsDuration()))
		}
	}
	// middleware filter
	opts = append(opts, http.Filter(
		CorsFilter(),  // 跨域 cors
		GetIpFilter(), // IP
	))
	// middleware
	opts = append(opts, http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(log.GetLogger()),
		HandlerMiddleware(),
		JwtMiddleware(),
	))
	// response
	opts = append(opts, http.ResponseEncoder(ResponseEncoder())) // http 正确-返回统一结构
	opts = append(opts, http.ErrorEncoder(ErrorEncoder()))       // http 错误-返回统一结构
	// http 服务
	srv := http.NewServer(opts...)
	// gin
	ginHttp := gin.New()
	ginHttp.GET("/swagger/*any", swaggerGin.WrapHandler(swaggerFiles.Handler))
	srv.HandlePrefix("/swagger/", ginHttp.Handler())
	// http 服务
	auth.RegisterSAuthHTTPServer(srv, service.NewAuthService())
	user.RegisterSUserHTTPServer(srv, service.NewUsersService())
	//
	return srv
}
