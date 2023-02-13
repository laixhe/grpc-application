package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "kratos-application/api/goproto"
	"kratos-application/app/goproto/internal/data"
)

type UserBiz struct {
	log *log.Helper
	db  *data.DB
}

func NewUserBiz(logger log.Logger, db *data.DB) *UserBiz {
	return &UserBiz{
		log: log.NewHelper(logger),
		db:  db,
	}
}

func (uc *UserBiz) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %v", req)
	return &pb.GetUserResponse{}, nil
}

func (uc *UserBiz) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.UserListResponse, error) {
	uc.log.WithContext(ctx).Infof("GetUserList: %v", req)
	return &pb.UserListResponse{}, nil
}
