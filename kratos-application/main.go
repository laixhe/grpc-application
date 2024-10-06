package main

import (
	"flag"
	"time"

	"github.com/go-kratos/kratos/v2"

	"kratos-application/app/server"
	"kratos-application/core/conf"
	"kratos-application/core/gormx"
	"kratos-application/core/i18nx"
	"kratos-application/core/logx"
	"kratos-application/core/redisx"
	"kratos-application/docs"
)

var (
	// Version 指定版本号 ( main.Version )
	Version string
)

var (
	// flagConfigFile 指定配置文件
	flagConfigFile string
)

func main() {
	// init config
	flag.Parse()
	flag.StringVar(&flagConfigFile, "config", "./config.yaml", "config path eg: -config config.yaml")
	conf.Init(flagConfigFile)
	// init api doc
	docs.SwaggerInfo.Description = docs.ErrorDescription()
	docs.SwaggerInfo.Version = Version + "-" + time.Now().Format(time.DateTime)
	conf.Get().GetApp().Version += "-" + Version
	// init log
	logx.Init(conf.Get().Log)
	// init other
	i18nx.Init()
	// init data db
	gormx.Init(conf.Get().GetDb())
	redisx.Init(conf.Get().GetRedis())
	//mongox.Init(conf.Get().GetMongodb())

	// init server
	httpServer := server.NewHTTPServer()
	kratosApp := kratos.New(
		kratos.Version(conf.Get().GetApp().Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			httpServer,
		),
	)
	if err := kratosApp.Run(); err != nil {
		panic(err)
	}
}
