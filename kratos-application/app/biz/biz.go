package biz

import (
	"github.com/go-kratos/kratos/v2/log"

	"kratos-application/app/data"
)

type Biz struct {
	log *log.Helper

	data *data.Data
}

func NewBiz(logger log.Logger) *Biz {
	return &Biz{
		log: log.NewHelper(logger),

		data: data.NewData(logger),
	}
}
