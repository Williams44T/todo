proto-gen:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ ./proto/service.proto