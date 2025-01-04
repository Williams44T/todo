package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient(ctx context.Context) (*DynamoDBClient, error) {
	defaultConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load default aws config: %v", err)
	}
	return &DynamoDBClient{
		client: dynamodb.NewFromConfig(defaultConfig),
	}, nil
}

func (ddb *DynamoDBClient) CreateTable(ctx context.Context) {
	ddb.client.CreateTable(ctx, &dynamodb.CreateTableInput{})
	ddb.client.GetItem(ctx, &dynamodb.GetItemInput{})
}
