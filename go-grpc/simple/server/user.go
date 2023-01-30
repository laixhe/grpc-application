package main

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	// 引入 proto 编译生成的包
	"github.com/laixhe/grpc-application/go-grpc/goproto"
)

// User 定义并实现约定的接口
type User struct {
	goproto.UnimplementedUserServer
}

// GetUser 获取某个 user 数据
func (u *User) GetUser(cxt context.Context,
	req *goproto.GetUserRequest) (*goproto.GetUserResponse, error) {

	if req.Userid <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "go-用户ID参数错误:%d", req.Userid)
	}

	// 待返回数据结构
	res := new(goproto.GetUserResponse)
	res.Userid = req.Userid
	res.Username = "go-laixhe"
	res.Sex = goproto.UserSex_MEN

	return res, nil
}

// GetUserList 获取 user 所有数据
func (u *User) GetUserList(cxt context.Context,
	req *goproto.GetUserListRequest) (*goproto.UserListResponse, error) {

	list := make([]*goproto.GetUserResponse, 0, 3)
	for i := 1000; i < 1004; i++ {
		list = append(list, &goproto.GetUserResponse{
			Userid:   int64(i),
			Username: "go-laixhe-" + strconv.Itoa(i),
			Sex:      goproto.UserSex_MEN,
		})
	}

	// 待返回数据结构
	res := new(goproto.UserListResponse)
	res.List = list

	return res, nil
}
