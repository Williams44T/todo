package service

import (
	"context"
	proto "todo/proto/gen/service"
)

type todoServer struct {
	proto.UnimplementedTodoServer
}

func NewTodoServer() *todoServer {
	return &todoServer{}
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
