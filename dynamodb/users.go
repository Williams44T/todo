package dynamodb

import (
	"context"
	"errors"
)

type AddUserReq struct{}
type AddUserResp struct{}

func (ddb *DynamoDBClient) AddUser(ctx context.Context, req *AddUserReq) (*AddUserResp, error) {
	return nil, errors.New("not implemented yet")
}

type GetUserReq struct{}
type GetUserResp struct{}

func (ddb *DynamoDBClient) GetUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	return nil, errors.New("not implemented yet")
}

type UpdateUserReq struct{}
type UpdateUserResp struct{}

func (ddb *DynamoDBClient) UpdateUser(ctx context.Context, req *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, errors.New("not implemented yet")
}

type DeleteUserReq struct{}
type DeleteUserResp struct{}

func (ddb *DynamoDBClient) DeleteUser(ctx context.Context, req *DeleteUserReq) (*DeleteUserResp, error) {
	return nil, errors.New("not implemented yet")
}
