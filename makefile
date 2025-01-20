proto-gen:
	protoc --go_out=./proto/ --go-grpc_out=./proto/ --proto_path=./proto ./proto/*.proto

server:
	go run ./cmd/service/service.go &

test:
	go test -v ./...

test-all:
	go test -v -tags integration ./...

build-cli:
	go build -o ./todo-cli ./cli/cli.go
