API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

.PHONY: api
api:
	rm -rf ./api/gen
	protoc --proto_path=./api \
		--proto_path=./third_party \
		--go_out=../ \
		--go-grpc_out=../ \
 		--grpc-gateway_out=../ \
		$(API_PROTO_FILES)
