package mock

import (
	"context"
	"errors"
	"todo/dynamodb"
)

type MockDynamoDBClient struct {
	// Users
	AddUserErr    error
	GetUserErr    error
	UpdateUserErr error
	DeleteUserErr error

	// Tasks
	AddTaskErr      error
	GetTaskErr      error
	BatchGetTaskErr error
	UpdateTaskErr   error
	DeleteTaskErr   error

	// Events
	AddEventErr      error
	GetEventErr      error
	BatchGetEventErr error
	UpdateEventErr   error
	DeleteEventErr   error
}

// assert that MockDynamoDBClient implements DynamoDBInterface
var _ dynamodb.DynamoDBInterface = &MockDynamoDBClient{}

func (mdb *MockDynamoDBClient) AddUser(ctx context.Context, req *dynamodb.AddUserReq) (*dynamodb.AddUserResp, error) {
	if mdb.AddUserErr != nil {
		return nil, mdb.AddUserErr
	}
	return &dynamodb.AddUserResp{}, nil
}

func (mdb *MockDynamoDBClient) GetUser(ctx context.Context, req *dynamodb.GetUserReq) (*dynamodb.GetUserResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) UpdateUser(ctx context.Context, req *dynamodb.UpdateUserReq) (*dynamodb.UpdateUserResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) DeleteUser(ctx context.Context, req *dynamodb.DeleteUserReq) (*dynamodb.DeleteUserResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) AddTask(ctx context.Context, req *dynamodb.AddTaskReq) (*dynamodb.AddTaskResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) GetTask(ctx context.Context, req *dynamodb.GetTaskReq) (*dynamodb.GetTaskResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) BatchGetTask(ctx context.Context, req *dynamodb.BatchGetTaskReq) (*dynamodb.BatchGetTaskResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) UpdateTask(ctx context.Context, req *dynamodb.UpdateTaskReq) (*dynamodb.UpdateTaskResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) DeleteTask(ctx context.Context, req *dynamodb.DeleteTaskReq) (*dynamodb.DeleteTaskResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) AddEvent(ctx context.Context, req *dynamodb.AddEventReq) (*dynamodb.AddEventResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) GetEvent(ctx context.Context, req *dynamodb.GetEventReq) (*dynamodb.GetEventResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) BatchGetEvent(ctx context.Context, req *dynamodb.BatchGetEventReq) (*dynamodb.BatchGetEventResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) UpdateEvent(ctx context.Context, req *dynamodb.UpdateEventReq) (*dynamodb.UpdateEventResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) DeleteEvent(ctx context.Context, req *dynamodb.DeleteEventReq) (*dynamodb.DeleteEventResp, error) {
	return nil, errors.New("not implemented")
}
