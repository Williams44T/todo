package service

import (
	"context"
	"fmt"
	"todo/dynamodb"
	proto "todo/proto/gen/service"
)

type todoServer struct {
	proto.UnimplementedTodoServer
	ddb *dynamodb.DynamoDBClient
}

func NewTodoServer(ctx context.Context) (*todoServer, error) {
	// get database client
	databaseClient, err := dynamodb.NewDynamoDBClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %v", err)
	}

	return &todoServer{
		ddb: databaseClient,
	}, nil
}

func (t *todoServer) GetTodo(ctx context.Context, req *proto.GetTodoReq) (*proto.GetTodoResp, error) {
	return &proto.GetTodoResp{
		Tasks: []*proto.Task{
			{
				Title:       "title",
				Description: "description",
			},
		},
	}, nil
}
