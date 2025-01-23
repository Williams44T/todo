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

func Test_TodoServer_GetAllTasks(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.GetAllTasksReq
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantTaskIdSet map[string]struct{}
		wantErr       bool
	}{
		{
			name: "happy path",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_1_ID: {
							{TaskID: common.TASK_1A_ID},
						},
						common.TEST_USER_2_ID: {
							{TaskID: common.TASK_2B_ID},
							{TaskID: common.TASK_2C_ID},
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				req: &proto.GetAllTasksReq{},
			},
			wantTaskIdSet: map[string]struct{}{
				common.TASK_2C_ID: {},
				common.TASK_2B_ID: {},
			},
			wantErr: false,
		},
		{
			name: "no user id in context",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_1_ID: {
							{TaskID: common.TASK_1A_ID},
						},
						common.TEST_USER_2_ID: {
							{TaskID: common.TASK_2B_ID},
							{TaskID: common.TASK_2C_ID},
						},
					},
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.GetAllTasksReq{},
			},
			wantTaskIdSet: map[string]struct{}{},
			wantErr:       true,
		},
		{
			name: "GetAllTasks returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					TasksTable: map[string][]dynamodb.Task{
						common.TEST_USER_1_ID: {
							{TaskID: common.TASK_1A_ID},
						},
						common.TEST_USER_2_ID: {
							{TaskID: common.TASK_2B_ID},
							{TaskID: common.TASK_2C_ID},
						},
					},
					GetAllTasksErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				req: &proto.GetAllTasksReq{},
			},
			wantTaskIdSet: map[string]struct{}{},
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TodoServer{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				ddb:                     tt.fields.ddb,
				jwt:                     tt.fields.jwt,
			}
			got, err := tr.GetAllTasks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoServer.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			gotTaskIdSet := map[string]struct{}{}
			for _, gotTask := range got.Tasks {
				gotTaskIdSet[gotTask.Id] = struct{}{}
			}
			for id := range tt.wantTaskIdSet {
				if _, ok := gotTaskIdSet[id]; !ok {
					t.Errorf("want task ID %v not returned", id)
				}
			}
			for id := range gotTaskIdSet {
				if _, ok := tt.wantTaskIdSet[id]; !ok {
					t.Errorf("unexpected task ID %v returned", id)
				}
			}
		})
	}
}
