package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2"

	"go-kratos/app/server"
	"go-kratos/core"
)

var (
	// Version 指定版本号 ( main.Version )
	Version string
	// ConfigFile 指定配置文件
	ConfigFile string
)

func init() {
	flag.StringVar(&ConfigFile, "conf", "./config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	logger := core.Logger()
	coreServer := core.NewServer(ConfigFile)
	httpServer := server.NewHTTPServer(coreServer, logger)
	kratosApp := kratos.New(kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(httpServer))
	if err := kratosApp.Run(); err != nil {
		panic(err)
	}
}
