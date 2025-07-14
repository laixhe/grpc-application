package controllers

import (
	"github.com/go-kratos/kratos/v2/log"

	"go-kratos/app/services"
	"go-kratos/core"
)

// Controller 业务控制器
type Controller struct {
	Auth *Auth
	User *User
}

func NewController(server *core.Server, logger log.Logger) *Controller {
	service := services.NewService(server, logger)
	return &Controller{
		Auth: newAuth(server, service, logger),
		User: newUser(server, service, logger),
	}
}
