package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"todo/common"
	"todo/dynamodb"
	ddbMock "todo/dynamodb/mock"
	proto "todo/proto/gen/service"
	"todo/service/token_manager"
	tmMock "todo/service/token_manager/mock"
)

func Test_todoServer_Signup(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.SignupReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.SignupResp
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "password",
				},
			},
			want: &proto.SignupResp{
				AccessJWT: "token_id_1",
			},
			wantErr: false,
		},
		{
			name: "empty first name",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty email",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "",
					Password:  "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty password",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "AddUser returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					AddUserErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IssueToken returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{
					IssueTokenErr: errors.New("test error"),
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SignupReq{
					FirstName: "Travis",
					LastName:  "Williams",
					Email:     "Williams44T@gmail.com",
					Password:  "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServer{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				ddb:                     tt.fields.ddb,
				jwt:                     tt.fields.jwt,
			}
			got, err := tr.Signup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// We can't guess the user id, so just check for existence and then erase it so the next check passes.
			if tt.want != nil && got != nil {
				if got.UserID == "" {
					t.Error("todoServer.Signup() no user id returned")
				}
				got.UserID = ""
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServer.Signup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_todoServer_Signin(t *testing.T) {
	hashedPassword, err := hashPassword(common.TEST_USER_1_PASSWORD)
	if err != nil {
		t.Errorf("todoServer.Signin() failed to hash password: %v", err)
	}
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.SigninReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.SigninResp
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UsersTable: map[string]dynamodb.User{
						common.TEST_USER_1_ID: {
							HashedPassword: hashedPassword,
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   common.TEST_USER_1_ID,
					Password: common.TEST_USER_1_PASSWORD,
				},
			},
			want: &proto.SigninResp{
				AccessJWT: "token_id_1",
			},
			wantErr: false,
		},
		{
			name: "empty user id",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UsersTable: map[string]dynamodb.User{
						common.TEST_USER_1_ID: {
							HashedPassword: hashedPassword,
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   "",
					Password: common.TEST_USER_1_PASSWORD,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid password",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UsersTable: map[string]dynamodb.User{
						common.TEST_USER_1_ID: {
							HashedPassword: hashedPassword,
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   common.TEST_USER_1_ID,
					Password: "invalid_password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "GetUser returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UsersTable: map[string]dynamodb.User{
						common.TEST_USER_1_ID: {
							HashedPassword: hashedPassword,
						},
					},
					GetUserErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   common.TEST_USER_1_ID,
					Password: common.TEST_USER_1_PASSWORD,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IssueToken returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UsersTable: map[string]dynamodb.User{
						common.TEST_USER_1_ID: {
							HashedPassword: hashedPassword,
						},
					},
				},
				jwt: &tmMock.MockTokenManager{
					IssueTokenErr: errors.New("test error"),
				},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.SigninReq{
					UserID:   common.TEST_USER_1_ID,
					Password: common.TEST_USER_1_PASSWORD,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &todoServer{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				ddb:                     tt.fields.ddb,
				jwt:                     tt.fields.jwt,
			}
			got, err := tr.Signin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServer.Signin() = %v, want %v", got, tt.want)
			}
		})
	}
}
