package services

import (
	"go-kratos/app/model/dao"
	"go-kratos/core"
)

// Service 业务服务逻辑
type Service struct {
	Auth *Auth
	User *User
}

func NewService(server *core.Server) *Service {
	modelDao := dao.NewDao(server)
	service := &Service{
		Auth: NewAuth(server, modelDao),
		User: NewUser(server, modelDao),
	}
	return service
}
