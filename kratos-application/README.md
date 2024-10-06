
#### 在 windows 系统使用 make
- 要 git、mingw64 配合使用
- 安装 git 
- 安装 mingw64 后复制 bin/mingw32-make.exe 为 /bin/make.exe
- 请使用 Git Bash 终端下执行 make 命令

- 用于生成 go protobuf struct 代码：
- --go_out
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

- 用于生成 grpc 服务代码：
- --go-grpc_out
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- 用于生成 http rest 服务代码：
- --go-http_out
```
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
```

- 用于给 protobuf 添加结构体 tag(可选)：
- protoc-go-inject-tag -input="./api/gen/v1api/*.pb.go"
```
go install github.com/favadi/protoc-go-inject-tag@latest
```

- 将 Go 的注释转换为 Swagger2.0 文档
- 在包含 main.go 文件的项目根目录运行 swag init 
```
go install github.com/swaggo/swag/cmd/swag@latest
```

- 用于生成 kratos 的错误定义代码(可选)：
- --go-errors_out
```
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
```

- 用于生成消息验证器代码(可选)：
```
go install github.com/envoyproxy/protoc-gen-validate@latest
```

- 用于生成 OpenAPI V3 文档(可选)：
- --openapi_out=fq_schema_naming=true,default_response=false:.
```
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
```
