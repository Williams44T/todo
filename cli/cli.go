package main

import (
	"log"
	"todo/cli/interceptor"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// get interceptors
	interceptor, err := interceptor.NewInterceptor()
	if err != nil {
		log.Fatalf("failed to get interceptors: %s", err)
	}

	// create client
	conn, err := grpc.NewClient(
		":9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.UnaryAuthMiddleware),
	)
	if err != nil {
		log.Fatalf("failed to create client conn: %s", err)
	}
	defer conn.Close()
	_ = proto.NewTodoClient(conn)
}
