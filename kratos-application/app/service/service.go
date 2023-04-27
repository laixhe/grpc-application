package service

import (
	"github.com/go-kratos/kratos/v2/log"

	"kratos-application/api/gen/v1api"
	"kratos-application/app/biz"
)

type V1Service struct {
	v1api.UnimplementedV1Server
	log *log.Helper

	biz *biz.Biz
}

func NewV1Service(logger log.Logger) *V1Service {
	return &V1Service{
		log: log.NewHelper(logger),

		biz: biz.NewBiz(logger),
	}
}
