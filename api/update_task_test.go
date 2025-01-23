package api

import (
	"context"
	"errors"
	"testing"
	"todo/common"
	"todo/interfaces/dynamodb"
	ddbMock "todo/interfaces/dynamodb/mock"
	"todo/interfaces/token_manager"
	tmMock "todo/interfaces/token_manager/mock"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func Test_todoServer_UpdateTask(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.UpdateTaskReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.UpdateTaskResp
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{
						Id: common.TASK_1A_ID,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "no task id",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{},
				},
			},
			wantErr: true,
		},
		{
			name: "no user id in context",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{
						Id: common.TASK_1A_ID,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "UpdateTask returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					UpdateTaskErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{
						Id: common.TASK_1A_ID,
					},
				},
			},
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
			_, err := tr.UpdateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
