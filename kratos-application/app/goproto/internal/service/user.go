package service

import (
	"context"

	pb "kratos-application/api/goproto"

	ubiz "kratos-application/app/goproto/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer

	userBiz *ubiz.UserBiz
}

func NewUserService(userBiz *ubiz.UserBiz) *UserService {
	return &UserService{
		userBiz: userBiz,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}

func (s *UserService) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.UserListResponse, error) {
	return &pb.UserListResponse{}, nil
}
