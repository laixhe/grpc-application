package biz

import (
	"kratos-application/app/cache"
	"kratos-application/app/data"
)

type Biz struct {
	Data  *data.Data
	Cache *cache.Cache
}

var newBiz = NewBiz()

func NewBiz() *Biz {
	return &Biz{
		Data:  data.NewData(),
		Cache: cache.NewCache(),
	}
}
