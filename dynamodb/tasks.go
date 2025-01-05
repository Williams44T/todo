package dynamodb

import (
	"context"
	"errors"
)

type AddTaskReq struct{}
type AddTaskResp struct{}

func (ddb *DynamoDBClient) AddTask(ctx context.Context, req *AddTaskReq) (*AddTaskResp, error) {
	return nil, errors.New("not implemented yet")
}

type GetTaskReq struct{}
type GetTaskResp struct{}

func (ddb *DynamoDBClient) GetTask(ctx context.Context, req *GetTaskReq) (*GetTaskResp, error) {
	return nil, errors.New("not implemented yet")
}

type BatchGetTaskReq struct{}
type BatchGetTaskResp struct{}

func (ddb *DynamoDBClient) BatchGetTask(ctx context.Context, req *BatchGetTaskReq) (*BatchGetTaskResp, error) {
	return nil, errors.New("not implemented yet")
}

type UpdateTaskReq struct{}
type UpdateTaskResp struct{}

func (ddb *DynamoDBClient) UpdateTask(ctx context.Context, req *UpdateTaskReq) (*UpdateTaskResp, error) {
	return nil, errors.New("not implemented yet")
}

type DeleteTaskReq struct{}
type DeleteTaskResp struct{}

func (ddb *DynamoDBClient) DeleteTask(ctx context.Context, req *DeleteTaskReq) (*DeleteTaskResp, error) {
	return nil, errors.New("not implemented yet")
}
