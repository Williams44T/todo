//go:build integration
// +build integration

package api

import (
	"context"
	"os"
	"testing"
	"todo/common"
	"todo/interfaces/dynamodb"
	"todo/interfaces/token_manager"
	proto "todo/proto/gen/go/api"
)

func Test_Integration_todoServer_Signup(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.SignupReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.Signup() failed to get database client: %v", err)
			}

			// get token manager
			jwtSecret, ok := os.LookupEnv(common.JWT_SECRET_ENV_VAR)
			if !ok {
				t.Errorf("%s must be provided as an environment variable", common.JWT_SECRET_ENV_VAR)
			}
			tokenManager, err := token_manager.NewTokenManager(jwtSecret)
			if err != nil {
				t.Errorf("failed to get token manager: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
				jwt: tokenManager,
			}

			_, err = tr.Signup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Integration_todoServer_Signin(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.SigninReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   common.TEST_USER_1_ID,
					Password: "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.Signin() failed to get database client: %v", err)
			}

			// get token manager
			jwtSecret, ok := os.LookupEnv(common.JWT_SECRET_ENV_VAR)
			if !ok {
				t.Errorf("%s must be provided as an environment variable", common.JWT_SECRET_ENV_VAR)
			}
			tokenManager, err := token_manager.NewTokenManager(jwtSecret)
			if err != nil {
				t.Errorf("failed to get token manager: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
				jwt: tokenManager,
			}

			_, err = tr.Signin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
