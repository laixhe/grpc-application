GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: api
api:
	rm -rf ./api/gen
	protoc --proto_path=./api \
		--proto_path=./third_party \
		--go_out=../ \
		--go-http_out=../ \
		--go-grpc_out=../ \
		$(API_PROTO_FILES)
	protoc-go-inject-tag -input="./api/gen/v1api/*.pb.go"
	swag init

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)"

