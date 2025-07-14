package services

import (
	"github.com/go-kratos/kratos/v2/log"

	"go-kratos/app/model/dao"
	"go-kratos/core"
)

// Service 业务服务逻辑
type Service struct {
	Auth *Auth
	User *User
}

func NewService(server *core.Server, logger log.Logger) *Service {
	modelDao := dao.NewDao(server)
	service := &Service{
		Auth: NewAuth(server, modelDao, logger),
		User: NewUser(server, modelDao, logger),
	}
	return service
}
