build: build-server build-client

build-server:
	GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o bin/hello_server hello_server/main.go
build-client:
	GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o bin/hello_client hello_client/main.go

gen-proto: ## Generate gRPC models
	 protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go_out=. helloworld.proto