package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/go-kratos/kratos/v2"

	"go-kratos/app/server"
	"go-kratos/core"
)

var (
	// GitVersion 指定版本号 ( go build -ldflags "-X main.GitVersion=xxx" )
	GitVersion string
	// ConfigFile 指定配置文件
	ConfigFile string
)

func init() {
	flag.StringVar(&ConfigFile, "conf", "./config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	fmt.Printf("[go version: %s] [config file: %s] [git: %s] \n", runtime.Version(), ConfigFile, GitVersion)

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
