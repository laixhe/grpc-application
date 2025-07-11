package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"

	"go-kratos/api/gen/pbauth"
	"go-kratos/api/gen/pbuser"
	"go-kratos/app/controllers"
	"go-kratos/app/server/middlewares"
	"go-kratos/core"
)

func NewHTTPServer(coreServer *core.Server) *kratosHttp.Server {
	conf := coreServer.Config()
	var opts = make([]kratosHttp.ServerOption, 0, 10)
	opts = append(opts, kratosHttp.Address(conf.Http.Addr()))
	if conf.Http.Timeout > 0 {
		opts = append(opts, kratosHttp.Timeout(time.Duration(conf.Http.Timeout)*time.Second))
	}
	// middleware filter
	opts = append(opts, kratosHttp.Filter(
		middlewares.CorsFilter(),
	))
	// middleware
	opts = append(opts, kratosHttp.Middleware(
		recovery.Recovery(),
	))

	controller := controllers.NewController(coreServer)
	srv := kratosHttp.NewServer(opts...)
	pbauth.RegisterSAuthHTTPServer(srv, controller.Auth)
	pbuser.RegisterSUserHTTPServer(srv, controller.User)

	return srv
}
