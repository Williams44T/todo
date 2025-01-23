package api

import (
	"context"
	"fmt"
	"os"
	"todo/common"
	"todo/interfaces/dynamodb"
	"todo/interfaces/token_manager"
	proto "todo/proto/gen/go/api"
)

type TodoServer struct {
	proto.UnimplementedTodoServer
	ddb dynamodb.DynamoDBInterface
	jwt token_manager.TokenManagerInterface
}

func NewTodoServer(ctx context.Context) (*TodoServer, error) {
	// get database client
	databaseClient, err := dynamodb.NewDynamoDBClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database client: %v", err)
	}

	// get token manager
	jwtSecret, ok := os.LookupEnv(common.JWT_SECRET_ENV_VAR)
	if !ok {
		return nil, fmt.Errorf("%s must be provided as an environment variable", common.JWT_SECRET_ENV_VAR)
	}
	tokenManager, err := token_manager.NewTokenManager(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get token manager: %v", err)
	}

	return &TodoServer{
		ddb: databaseClient,
		jwt: tokenManager,
	}, nil
}
