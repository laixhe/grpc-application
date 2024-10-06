#### 生成工具
- protoc                  是 protobuf 文件编译器
- protoc-gen-go           插件生成 protobuf go 代码
- protoc-gen-go-grpc      插件生成 gPRC 代码(服务端、客户端)
- protoc-gen-grpc-gateway 插件生成 grpc http 网关服务
- protoc-gen-openapiv2    插件生成 Swagger 文档
```
# protoc
https://github.com/protocolbuffers/protobuf/releases
protoc --version

# protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
protoc-gen-go --version

# protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc-gen-go-grpc --version

# protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
protoc-gen-grpc-gateway --version

# protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
protoc-gen-openapiv2 --version
```

#### proto 文件编译
```
protoc \
--go_opt=paths=source_relative --go-grpc_out=. \
--go-grpc_opt=paths=source_relative --go_out=. \
*.proto

protoc 参数说明
--proto_path     指定要在其中搜索导入（import）的 proto 文件目录(简写 -I 如果没有指定参数，则在当前目录进行搜索)
--go-grpc_out=   生成的 go-grpc 源码保存目录
--go_out         生成的 go 源码保存目录
```

#### proto 文件编译 getaway logtostderr=true:
```
protoc \
--go_opt=paths=source_relative --go-grpc_out=. \
--go-grpc_opt=paths=source_relative --go_out=. \
--grpc-gateway_opt=paths=source_relative --grpc-gateway_out=. \
*.proto
```
