package controllers

import (
	"github.com/go-kratos/kratos/v2/log"

	"go-kratos/api/gen/pbuser"
	"go-kratos/app/services"
	"go-kratos/core"
)

type User struct {
	pbuser.UnimplementedSUserServer
	log     *log.Helper
	server  *core.Server
	service *services.Service
}

func newUser(server *core.Server, service *services.Service, logger log.Logger) *User {
	return &User{
		log:     log.NewHelper(logger),
		server:  server,
		service: service,
	}
}
