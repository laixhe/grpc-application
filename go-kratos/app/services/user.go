package services

import (
	"go-kratos/app/model/dao"
	"go-kratos/core"
)

type User struct {
	server *core.Server
	dao    *dao.Dao
}

func NewUser(server *core.Server, modelDao *dao.Dao) *User {
	return &User{
		server: server,
		dao:    modelDao,
	}
}
