proto-gen:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ ./proto/service.proto

server:
	go run ./cmd/service/service.go &

test:
	go test -tags=\!integration ./...

build-cli:
	go build -o ./todo-cli ./cli/cli.go
