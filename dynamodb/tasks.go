package dynamodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type RecurringRule struct {
	CronExpression string `dynamodbav:"cron_expression"`
	StartDate      int64  `dynamodbav:"start_date"`
	EndDate        int64  `dynamodbav:"end_date"`
}

type Task struct {
	ID            string        `dynamodbav:"id"`
	UserID        string        `dynamodbav:"user_id"`
	Title         string        `dynamodbav:"title"`
	Description   string        `dynamodbav:"description"`
	Status        string        `dynamodbav:"status"`
	Tags          []string      `dynamodbav:"tags"`
	Parents       []string      `dynamodbav:"parents"`
	DueDate       int64         `dynamodbav:"due_date"`
	RecurringRule RecurringRule `dynamodbav:"recurring_rule"`
}

type AddTaskReq struct {
	Task Task
}
type AddTaskResp struct{}

func (ddb *DynamoDBClient) AddTask(ctx context.Context, req *AddTaskReq) (*AddTaskResp, error) {
	item, err := attributevalue.MarshalMap(req.Task)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal task: %v", err)
	}
	_, err = ddb.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &ddb.tasksTableName,
		Item:      item,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to put user into users table: %v", err)
	}
	return &AddTaskResp{}, nil
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
