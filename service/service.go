package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"todo/dynamodb"
	proto "todo/proto/gen/service"
	"todo/service/token_manager"

	"github.com/google/uuid"
)

const (
	JWT_SECRET_ENV_KEY = "JWT_SECRET"
)

type todoServer struct {
	proto.UnimplementedTodoServer
	ddb dynamodb.DynamoDBInterface
	jwt token_manager.TokenManagerInterface
}

func NewTodoServer(ctx context.Context) (*todoServer, error) {
	// get database client
	databaseClient, err := dynamodb.NewDynamoDBClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %v", err)
	}

	// get token manager
	jwtSecret, ok := os.LookupEnv(JWT_SECRET_ENV_KEY)
	if !ok {
		return nil, fmt.Errorf("%s must be provided as an environment variable", JWT_SECRET_ENV_KEY)
	}
	tokenManager, err := token_manager.NewTokenManager(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get token manager: %v", err)
	}

	return &todoServer{
		ddb: databaseClient,
		jwt: tokenManager,
	}, nil
}

// TODO: implement this
func hashPassword(password string) (string, error) {
	return password, nil
}

func (t *todoServer) Signup(ctx context.Context, req *proto.SignupReq) (*proto.SignupResp, error) {
	// validate request
	if req.FirstName == "" {
		return nil, errors.New("first name cannot be blank")
	}
	// TODO: implement stronger email validation
	if req.Email == "" {
		return nil, errors.New("email cannot be blank")
	}
	// TODO: implement stronger password validation
	if req.Password == "" {
		return nil, errors.New("password cannot be blank")
	}

	// hash password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// generate user id
	userID := uuid.New()

	// add user to DDB
	_, err = t.ddb.AddUser(ctx, &dynamodb.AddUserReq{
		User: dynamodb.User{
			ID:             hashedPassword,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			HashedPassword: hashedPassword,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add user: %v", err)
	}

	// generate access token
	token, err := t.jwt.IssueToken(userID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to issue jwt: %v", err)
	}

	// TODO: generate refresh token

	return &proto.SignupResp{
		AccessJWT: token,
	}, nil
}

func (t *todoServer) GetTodo(ctx context.Context, req *proto.GetTodoReq) (*proto.GetTodoResp, error) {
	return &proto.GetTodoResp{
		Tasks: []*proto.Task{
			{
				Title:       "title",
				Description: "description",
			},
		},
	}, nil
}
