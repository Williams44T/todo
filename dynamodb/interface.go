package dynamodb

import "context"

type DynamoDBInterface interface {
	// Users
	AddUser(context.Context, *AddUserReq) (*AddUserResp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
	DeleteUser(context.Context, *DeleteUserReq) (*DeleteUserResp, error)

	// Tasks
	AddTask(context.Context, *AddTaskReq) (*AddTaskResp, error)
	GetTask(context.Context, *GetTaskReq) (*GetTaskResp, error)
	BatchGetTask(context.Context, *BatchGetTaskReq) (*BatchGetTaskResp, error)
	UpdateTask(context.Context, *UpdateTaskReq) (*UpdateTaskResp, error)
	DeleteTask(context.Context, *DeleteTaskReq) (*DeleteTaskResp, error)

	// Events
	AddEvent(context.Context, *AddEventReq) (*AddEventResp, error)
	GetEvent(context.Context, *GetEventReq) (*GetEventResp, error)
	BatchGetEvent(context.Context, *BatchGetEventReq) (*BatchGetEventResp, error)
	UpdateEvent(context.Context, *UpdateEventReq) (*UpdateEventResp, error)
	DeleteEvent(context.Context, *DeleteEventReq) (*DeleteEventResp, error)
}
