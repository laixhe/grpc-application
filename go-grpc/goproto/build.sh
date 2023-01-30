rm *.go

protoc --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. *.proto