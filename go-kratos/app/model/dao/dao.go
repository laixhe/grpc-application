package dao

import (
	"context"

	"github.com/laixhe/gonet/orm/orm"
	"gorm.io/gorm"

	"go-kratos/app/model"
	"go-kratos/core"
)

// Dao 业务数据操作
type Dao struct {
	server     *core.Server
	User       *model.User
	UserExtend *model.UserExtend
}

func NewDao(server *core.Server) *Dao {
	return &Dao{
		server:     server,
		User:       &model.User{},
		UserExtend: &model.UserExtend{},
	}
}

func (d *Dao) Orm() orm.Client {
	return d.server.Orm()
}

func (d *Dao) WithContext(ctx context.Context) *gorm.DB {
	return d.server.Orm().WithContext(ctx)
}
