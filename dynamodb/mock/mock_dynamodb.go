package mock

import (
	"context"
	"errors"
	"todo/dynamodb"
)

type MockDynamoDBClient struct {
	// Tables
	UsersTable map[string]dynamodb.User
	TasksTable map[string]dynamodb.Task

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
	if mdb.UsersTable == nil {
		mdb.UsersTable = make(map[string]dynamodb.User)
	}
	mdb.UsersTable[req.User.ID] = req.User
	return &dynamodb.AddUserResp{}, nil
}

func (mdb *MockDynamoDBClient) GetUser(ctx context.Context, req *dynamodb.GetUserReq) (*dynamodb.GetUserResp, error) {
	if mdb.GetUserErr != nil {
		return nil, mdb.GetUserErr
	}
	if mdb.UsersTable == nil {
		return nil, errors.New("UsersTable does not exist")
	}
	user := mdb.UsersTable[req.ID]
	return &dynamodb.GetUserResp{
		User: &user,
	}, nil
}

func (mdb *MockDynamoDBClient) UpdateUser(ctx context.Context, req *dynamodb.UpdateUserReq) (*dynamodb.UpdateUserResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) DeleteUser(ctx context.Context, req *dynamodb.DeleteUserReq) (*dynamodb.DeleteUserResp, error) {
	return nil, errors.New("not implemented")
}

func (mdb *MockDynamoDBClient) AddTask(ctx context.Context, req *dynamodb.AddTaskReq) (*dynamodb.AddTaskResp, error) {
	if mdb.AddTaskErr != nil {
		return nil, mdb.AddTaskErr
	}
	if mdb.TasksTable == nil {
		mdb.TasksTable = make(map[string]dynamodb.Task)
	}
	mdb.TasksTable[req.Task.TaskID] = req.Task
	return &dynamodb.AddTaskResp{}, nil
}

func (mdb *MockDynamoDBClient) GetTask(ctx context.Context, req *dynamodb.GetTaskReq) (*dynamodb.GetTaskResp, error) {
	if mdb.GetTaskErr != nil {
		return nil, mdb.GetTaskErr
	}
	if mdb.TasksTable == nil {
		return nil, errors.New("tasksTable does not exist")
	}
	task := mdb.TasksTable[req.TaskID]
	if task.UserID != req.UserID {
		return nil, errors.New("unauthorized")
	}
	return &dynamodb.GetTaskResp{
		Task: &task,
	}, nil
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
