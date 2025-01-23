package main

import (
	"context"
	"log"
	"net"
	proto "todo/proto/gen/go/api"
	"todo/api"
	"todo/api/interceptor"

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

	// get interceptor
	interceptor, err := interceptor.NewInterceptor()
	if err != nil {
		log.Fatalf("failed to get interceptor: %s", err)
	}

	// create server
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.UnaryAuthMiddleware))

	// register server
	todoService, err := api.NewTodoServer(ctx)
	if err != nil {
		log.Fatalf("failed to get todo api. %s", err)
	}
	proto.RegisterTodoServer(server, todoService)

	// egister reflection api.on server
	reflection.Register(server)

	// start server
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
