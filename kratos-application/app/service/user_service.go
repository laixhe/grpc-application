package service

import (
	"kratos-application/api/gen/user"
	"kratos-application/app/biz"
)

type UsersService struct {
	user.UnimplementedSUserServer
	userBiz *biz.UserBiz
}

func NewUsersService() *UsersService {
	return &UsersService{
		userBiz: biz.NewUserBiz(),
	}
}
