package dynamodb

import (
	"context"
	"errors"
)

type AddEventReq struct{}
type AddEventResp struct{}

func (ddb *DynamoDBClient) AddEvent(ctx context.Context, req *AddEventReq) (*AddEventResp, error) {
	return nil, errors.New("not implemented yet")
}

type GetEventReq struct{}
type GetEventResp struct{}

func (ddb *DynamoDBClient) GetEvent(ctx context.Context, req *GetEventReq) (*GetEventResp, error) {
	return nil, errors.New("not implemented yet")
}

type BatchGetEventReq struct{}
type BatchGetEventResp struct{}

func (ddb *DynamoDBClient) BatchGetEvent(ctx context.Context, req *BatchGetEventReq) (*BatchGetEventResp, error) {
	return nil, errors.New("not implemented yet")
}

type UpdateEventReq struct{}
type UpdateEventResp struct{}

func (ddb *DynamoDBClient) UpdateEvent(ctx context.Context, req *UpdateEventReq) (*UpdateEventResp, error) {
	return nil, errors.New("not implemented yet")
}

type DeleteEventReq struct{}
type DeleteEventResp struct{}

func (ddb *DynamoDBClient) DeleteEvent(ctx context.Context, req *DeleteEventReq) (*DeleteEventResp, error) {
	return nil, errors.New("not implemented yet")
}
