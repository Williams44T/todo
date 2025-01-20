package dynamodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

// AddUser puts a user into the users table exactly as given.
func (ddb *DynamoDBClient) AddUser(ctx context.Context, req *AddUserReq) (*AddUserResp, error) {
	item, err := attributevalue.MarshalMap(req.User)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %v", err)
	}
	_, err = ddb.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &ddb.usersTableName,
		Item:      item,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to put user into users table: %v", err)
	}
	return &AddUserResp{}, nil
}

type GetUserReq struct {
	ID string `dynamodbav:"id"`
}
type GetUserResp struct {
	User *User
}

// GetUser uses the given user id to find a user.
// User will be nil if no user is found.
func (ddb *DynamoDBClient) GetUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	key, err := attributevalue.MarshalMap(req.ID)
	fmt.Println(key)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user ID: %v", err)
	}
	getItemResp, err := ddb.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &ddb.usersTableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: req.ID},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	var user *User
	if getItemResp.Item != nil {
		err = attributevalue.UnmarshalMap(getItemResp.Item, user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal user: %v", err)
		}
	}
	return &GetUserResp{
		User: user,
	}, nil
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
