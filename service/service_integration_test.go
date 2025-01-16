package service

import (
	"context"
	"os"
	"testing"
	"todo/dynamodb"
	proto "todo/proto/gen/service"
	"todo/service/token_manager"
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
			jwtSecret, ok := os.LookupEnv(JWT_SECRET_ENV_KEY)
			if !ok {
				t.Errorf("%s must be provided as an environment variable", JWT_SECRET_ENV_KEY)
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
