//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-application/app/goproto/internal/biz"
	"kratos-application/app/goproto/internal/conf"
	"kratos-application/app/goproto/internal/data"
	"kratos-application/app/goproto/internal/server"
	"kratos-application/app/goproto/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
