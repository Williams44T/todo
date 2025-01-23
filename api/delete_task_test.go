package api

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"todo/common"
	"todo/interfaces/dynamodb"
	ddbMock "todo/interfaces/dynamodb/mock"
	"todo/interfaces/token_manager"
	tmMock "todo/interfaces/token_manager/mock"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func Test_todoServer_DeleteTask(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.DeleteTaskReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.DeleteTaskResp
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
				req: &proto.DeleteTaskReq{
					TaskId: "task_id",
				},
			},
			want:    &proto.DeleteTaskResp{},
			wantErr: false,
		},
		{
			name: "no task id provided",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.DeleteTaskReq{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no user id provided in context",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.DeleteTaskReq{
					TaskId: "task_id",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "DDB DeleteTask throws error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					DeleteTaskErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.DeleteTaskReq{
					TaskId: "task_id",
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
			got, err := tr.DeleteTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServer.DeleteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
