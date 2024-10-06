package comm

import (
	"context"
	"encoding/json"
	"net/http"

	// 引入 proto 编译生成的包
	"go-grpc/api/gen/v1api"
)

var V1Client v1api.V1Client

// AuthRegister 注册用户
func AuthRegister(w http.ResponseWriter, r *http.Request) {

	// 获取 Post 的参数
	email := r.FormValue("email")
	password := r.FormValue("password")

	// 调用 Grpc 的远程接口
	data, err := V1Client.AuthRegister(context.Background(), &v1api.AuthRegisterRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		_, _ = w.Write([]byte("error: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)
	_, _ = w.Write(js)
}

// AuthLogin 登录
func AuthLogin(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	// 调用 Grpc 的远程接口
	data, err := V1Client.AuthLogin(context.Background(), &v1api.AuthLoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		_, _ = w.Write([]byte("error: " + err.Error()))
		return
	}

	// json 格式化
	js, _ := json.Marshal(data)
	_, _ = w.Write(js)
}
