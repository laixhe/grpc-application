package controllers

import (
	"go-kratos/app/services"
	"go-kratos/core"
)

// Controller 业务控制器
type Controller struct {
	Auth *Auth
	User *User
}

func NewController(server *core.Server) *Controller {
	service := services.NewService(server)
	return &Controller{
		Auth: newAuth(server, service),
		User: newUser(server, service),
	}
}
