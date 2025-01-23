package api

import (
	"context"
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

func Test_todoServer_GetTask(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.GetTaskReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.GetTaskResp
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_1_ID: {
							{
								TaskID: "task_id",
								Title:  "task title",
								Status: proto.Status_INCOMPLETE.String(),
							},
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.GetTaskReq{
					Id: "task_id",
				},
			},
			want: &proto.GetTaskResp{
				Task: &proto.Task{
					Id:     "task_id",
					Title:  "task title",
					Status: proto.Status_INCOMPLETE,
				},
			},
			wantErr: false,
		},
		{
			name: "no user id in context",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_1_ID: {
							{
								TaskID: "task_id",
								Title:  "task title",
								Status: proto.Status_INCOMPLETE.String(),
							},
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetTaskReq{
					Id: "task_id",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "mismatched user id",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_2_ID: {
							{
								TaskID: "task_id",
								Title:  "task title",
								Status: proto.Status_INCOMPLETE.String(),
							},
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.GetTaskReq{
					Id: "task_id",
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
			got, err := tr.GetTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServer.GetTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
