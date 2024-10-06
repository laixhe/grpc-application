package biz

import (
	"kratos-application/app/cache"
	"kratos-application/app/data"
)

type UserBiz struct {
	Data  *data.Data
	Cache *cache.Cache
}

func NewUserBiz() *UserBiz {
	return &UserBiz{
		Data:  newBiz.Data,
		Cache: newBiz.Cache,
	}
}
