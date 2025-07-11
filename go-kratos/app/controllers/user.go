package controllers

import (
	"go-kratos/api/gen/pbuser"
	"go-kratos/app/services"
	"go-kratos/core"
)

type User struct {
	pbuser.UnimplementedSUserServer
	server  *core.Server
	service *services.Service
}

func newUser(server *core.Server, service *services.Service) *User {
	return &User{
		server:  server,
		service: service,
	}
}
