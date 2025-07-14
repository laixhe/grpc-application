package services

import (
	"github.com/go-kratos/kratos/v2/log"

	"go-kratos/app/model/dao"
	"go-kratos/core"
)

type User struct {
	log    *log.Helper
	server *core.Server
	dao    *dao.Dao
}

func NewUser(server *core.Server, modelDao *dao.Dao, logger log.Logger) *User {
	return &User{
		log:    log.NewHelper(logger),
		server: server,
		dao:    modelDao,
	}
}
