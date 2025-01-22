package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	UserIDKey        = "user_id"
	TaskIDKey        = "task_id"
	TitleKey         = "title"
	DescriptionKey   = "description"
	StatusKey        = "status"
	TagsKey          = "tags"
	ParentsKey       = "parents"
	DueDateKey       = "due_date"
	RecurringRuleKey = "recurring_rule"
)

type RecurringRule struct {
	CronExpression string `dynamodbav:"cron_expression"`
	StartDate      int64  `dynamodbav:"start_date"`
	EndDate        int64  `dynamodbav:"end_date"`
}

type Task struct {
	UserID        string         `dynamodbav:"user_id"`
	TaskID        string         `dynamodbav:"task_id"`
	Title         string         `dynamodbav:"title"`
	Description   string         `dynamodbav:"description"`
	Status        string         `dynamodbav:"status"`
	Tags          []string       `dynamodbav:"tags"`
	Parents       []string       `dynamodbav:"parents"`
	DueDate       int64          `dynamodbav:"due_date"`
	RecurringRule *RecurringRule `dynamodbav:"recurring_rule"`
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
		return nil, fmt.Errorf("failed to put task into tasks table: %v", err)
	}
	return &AddTaskResp{}, nil
}

type GetTaskReq struct {
	UserID string
	TaskID string
}
type GetTaskResp struct {
	Task *Task
}

func (ddb *DynamoDBClient) GetTask(ctx context.Context, req *GetTaskReq) (*GetTaskResp, error) {
	getItemResp, err := ddb.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &ddb.tasksTableName,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: req.UserID},
			"task_id": &types.AttributeValueMemberS{Value: req.TaskID},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	var task *Task
	if getItemResp.Item != nil {
		task = &Task{}
		err = attributevalue.UnmarshalMap(getItemResp.Item, task)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal task: %v", err)
		}
	}
	return &GetTaskResp{
		Task: task,
	}, nil
}

type BatchGetTaskReq struct{}
type BatchGetTaskResp struct{}

func (ddb *DynamoDBClient) BatchGetTask(ctx context.Context, req *BatchGetTaskReq) (*BatchGetTaskResp, error) {
	return nil, errors.New("not implemented yet")
}

type GetAllTasksReq struct {
	UserID   string
	Statuses []string
}
type GetAllTasksResp struct {
	Tasks []Task
}

func (ddb *DynamoDBClient) GetAllTasks(ctx context.Context, req *GetAllTasksReq) (*GetAllTasksResp, error) {
	keyEx := expression.Key("user_id").Equal(expression.Value(req.UserID))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build expression: %v", err)
	}
	queryPaginator := dynamodb.NewQueryPaginator(ddb.client, &dynamodb.QueryInput{
		TableName:                 aws.String(ddb.tasksTableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	})
	var tasks []Task
	for queryPaginator.HasMorePages() {
		response, err := queryPaginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to query ddb: %v", err)
		} else {
			var taskPage []Task
			err = attributevalue.UnmarshalListOfMaps(response.Items, &taskPage)
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal query response: %v", err)
			} else {
				tasks = append(tasks, taskPage...)
			}
		}
	}
	return &GetAllTasksResp{
		Tasks: tasks,
	}, nil
}

type UpdateTaskReq struct {
	UserID  string
	TaskID  string
	KVPairs map[string]interface{}
}
type UpdateTaskResp struct {
	Task Task
}

func buildUpdateExpression(kvPairs map[string]interface{}) (*expression.UpdateBuilder, error) {
	update := expression.Set(expression.Name("updated_at"), expression.Value(time.Now().Unix()))
	for name, value := range kvPairs {
		switch name {
		case TitleKey, DescriptionKey, StatusKey:
			if _, ok := value.(string); !ok {
				return nil, fmt.Errorf("the value type of %s should be a string", name)
			}
		case DueDateKey:
			if _, ok := value.(int64); !ok {
				return nil, fmt.Errorf("the value type of %s should be int64", name)
			}
		case TagsKey, ParentsKey:
			if _, ok := value.([]string); !ok {
				return nil, fmt.Errorf("the value type of %s should be a list of strings", name)
			}
		case RecurringRuleKey:
			if value == nil {
				continue
			}
			if _, ok := value.(*RecurringRule); !ok {
				return nil, fmt.Errorf("the value type of %s should model the RecurringRule Type", name)
			}
		case UserIDKey, TaskIDKey:
			return nil, fmt.Errorf("not allowed to update %s", name)
		default:
			return nil, fmt.Errorf("unknown task attribute: %s", name)
		}
		update.Set(expression.Name(name), expression.Value(value))
	}
	return &update, nil
}

func (ddb *DynamoDBClient) UpdateTask(ctx context.Context, req *UpdateTaskReq) (*UpdateTaskResp, error) {
	update, err := buildUpdateExpression(req.KVPairs)
	if err != nil {
		return nil, fmt.Errorf("failed to get update builder: %v", err)
	}
	expr, err := expression.NewBuilder().WithUpdate(*update).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build expression: %v", err)
	}
	resp, err := ddb.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &ddb.tasksTableName,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: req.UserID},
			"task_id": &types.AttributeValueMemberS{Value: req.TaskID},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update item: %v", err)
	}
	updatedTask := Task{}
	if err = attributevalue.UnmarshalMap(resp.Attributes, &updatedTask); err != nil {
		return nil, fmt.Errorf("failed to unmarshal attribute map: %v", err)
	}
	return &UpdateTaskResp{
		Task: updatedTask,
	}, nil
}

type DeleteTaskReq struct{}
type DeleteTaskResp struct{}

func (ddb *DynamoDBClient) DeleteTask(ctx context.Context, req *DeleteTaskReq) (*DeleteTaskResp, error) {
	return nil, errors.New("not implemented yet")
}
