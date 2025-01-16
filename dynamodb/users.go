package dynamodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type User struct {
	ID             string `dynamodbav:"id"`
	FirstName      string `dynamodbav:"first_name"`
	LastName       string `dynamodbav:"last_name"`
	Email          string `dynamodbav:"email"`
	HashedPassword string `dynamodbav:"hashed_password"`
}

type AddUserReq struct {
	User User
}
type AddUserResp struct{}

func (ddb *DynamoDBClient) AddUser(ctx context.Context, req *AddUserReq) (*AddUserResp, error) {
	item, err := attributevalue.MarshalMap(req.User)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %v", err)
	}
	ddb.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &ddb.usersTableName,
		Item:      item,
	})
	return &AddUserResp{}, nil
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
