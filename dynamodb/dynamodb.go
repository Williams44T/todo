package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	client          *dynamodb.Client
	usersTableName  string
	tasksTableName  string
	eventsTableName string
}

// make client implement defined interface
var _ DynamoDBInterface = &DynamoDBClient{}

func NewDynamoDBClient(ctx context.Context) (*DynamoDBClient, error) {
	defaultConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default aws config: %v", err)
	}
	return &DynamoDBClient{
		client:          dynamodb.NewFromConfig(defaultConfig),
		usersTableName:  "todo-users",
		tasksTableName:  "todo-tasks",
		eventsTableName: "todo-events",
	}, nil
}
