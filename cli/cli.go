package main

import (
	"context"
	"fmt"
	"log"
	proto "todo/proto/gen/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create client
	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:9001", options)
	if err != nil {
		log.Fatalf("failed to create client conn: %s", err)
	}
	defer conn.Close()
	client := proto.NewTodoClient(conn)

	// GetTodo
	resp, err := client.GetTodo(context.Background(), &proto.GetTodoReq{})
	if err != nil {
		log.Fatalf("failed to GetTodo")
	}
	fmt.Print(resp)
}
