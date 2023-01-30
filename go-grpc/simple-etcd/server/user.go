package main

import (
	"context"
	"strconv"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/grpc-application/go-grpc/goproto"
)

// User 定义并实现约定的接口
type User struct {
	pb.UnimplementedUserServer
}

// GetUser 获取某个 user 数据
func (u *User) GetUser(cxt context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 待返回数据结构
	res := new(pb.GetUserResponse)
	res.Userid = req.Userid
	res.Username = "go-laixhe-etcd"
	res.Sex = pb.UserSex_MEN

	return res, nil
}

// GetUserList 获取 user 所有数据
func (u *User) GetUserList(cxt context.Context, req *pb.GetUserListRequest) (*pb.UserListResponse, error) {

	list := make([]*pb.GetUserResponse, 0, 3)
	for i := 1; i <= 3; i++ {
		list = append(list, &pb.GetUserResponse{
			Userid:   int64(i),
			Username: "go-laixhe-etcd-" + strconv.Itoa(i),
			Sex:      pb.UserSex_MEN,
		})
	}
	// 待返回数据结构
	resp := new(pb.UserListResponse)
	resp.List = list

	return resp, nil
}
