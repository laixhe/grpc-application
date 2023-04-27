package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"kratos-application/app/server"
	"kratos-application/core/config"
	"kratos-application/core/gormx"
	"kratos-application/core/redisx"
	"kratos-application/docs"
)

var (
	// Version 指定版本号 ( main.Version )
	Version string
)

// flagConfigFile 指定配置文件
var flagConfigFile string

// @title	API接口
func main() {
	// init config
	flag.Parse()
	flag.StringVar(&flagConfigFile, "config", "./config/config.yaml", "config path eg: -config config/config.yaml")
	config.Init(flagConfigFile)
	// init api doc
	docs.SwaggerInfo.Description = docs.ErrorDescription()
	if Version != "" {
		docs.SwaggerInfo.Version = Version
		config.Get().App.Version = Version
	}
	// init log
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	// init server
	grpcServer := server.NewGRPCServer(config.Get().Servers.Grpc, logger)
	httpServer := server.NewHTTPServer(config.Get().Servers.Http, config.Get().Jwt, logger)
	app := kratos.New(
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			grpcServer,
			httpServer,
		),
	)
	// init data db
	gormx.Init(config.Get().Mysql)
	redisx.Init(config.Get().Redis)
	// run
	if err := app.Run(); err != nil {
		panic(err)
	}
}
