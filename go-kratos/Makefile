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
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: api
api:
	rm -rf ./api/gen
	protoc --proto_path=./api \
		--proto_path=./third_party \
		--go_out=../ \
		--go-http_out=../ \
		--go-grpc_out=../ \
		--openapi_out=fq_schema_naming=true,default_response=false:. \
		$(API_PROTO_FILES)

.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)"

