package main

import (
	"context"
	"log"
	"net"
	proto "todo/proto/gen/service"
	"todo/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	// create listener; it's over 9000!
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}

	// create server
	server := grpc.NewServer()

	// register server
	todoService, err := service.NewTodoServer(ctx)
	if err != nil {
		log.Fatalf("failed to get todo service: %s", err)
	}
	proto.RegisterTodoServer(server, todoService)

	// egister reflection service on server
	reflection.Register(server)

	// start server
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
