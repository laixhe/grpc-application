package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	// 引入 proto 编译生成的包
	pb "github.com/laixhe/grpc-application/go-grpc/goproto"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	// 获取 GET 的参数
	userid := r.FormValue("userid")
	id, err := strconv.ParseInt(userid, 10, 0)
	if err != nil {
		w.Write([]byte("userid The parameters must be integers"))
		return
	}

	// 调用 Grpc 的远程接口
	data, err := UClient.GetUser(context.Background(), &pb.GetUserRequest{Userid: id})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)

	w.Write(js)

}

func GetUserList(w http.ResponseWriter, r *http.Request) {

	// 调用 Grpc 的远程接口
	data, err := UClient.GetUserList(context.Background(), &pb.GetUserListRequest{})
	if err != nil {
		w.Write([]byte("Grpc: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data.List)

	w.Write(js)

}
