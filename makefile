proto-gen:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ --proto_path=./proto ./proto/*.proto

server:
	go run ./cmd/service/service.go &

test-all: test test-integration

test:
	go test -v ./...

test-integration:
	go test -v -tags integration ./...

build-cli:
	go build -o ./todo-cli ./cli/cli.go
