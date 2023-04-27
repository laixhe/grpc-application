package server

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-application/api/gen/v1api"
	"kratos-application/app/service"
	"kratos-application/core/config"
)

func NewGRPCServer(c config.Server, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	opts = append(opts, grpc.Address(c.Addr()))
	if c.Timeout > 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Timeout)*time.Second))
	}

	// grpc 服务
	srv := grpc.NewServer(opts...)
	// v1 grpc 服务
	v1api.RegisterV1Server(srv, service.NewV1Service(logger))
	return srv
}
