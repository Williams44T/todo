package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
	"todo/common"
	"todo/dynamodb"
	ddbMock "todo/dynamodb/mock"
	proto "todo/proto/gen/service"
	"todo/service/token_manager"
	tmMock "todo/service/token_manager/mock"

	"google.golang.org/grpc/metadata"
)

func Test_todoServer_AddTask(t *testing.T) {
	type fields struct {
		UnimplementedTodoServer proto.UnimplementedTodoServer
		ddb                     dynamodb.DynamoDBInterface
		jwt                     token_manager.TokenManagerInterface
	}
	type args struct {
		ctx context.Context
		req *proto.AddTaskReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.AddTaskResp
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
				req: &proto.AddTaskReq{
					Title:       "do something",
					Description: "do something with extra steps",
					Tags:        []string{"tag1", "tag2"},
					DueDate:     time.Now().Unix(),
				},
			},
			want:    &proto.AddTaskResp{},
			wantErr: false,
		},
		{
			name: "no user id in context",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: context.Background(),
				req: &proto.AddTaskReq{
					Title:       "do something",
					Description: "do something with extra steps",
					Tags:        []string{"tag1", "tag2"},
					DueDate:     time.Now().Unix(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no title",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.AddTaskReq{
					Description: "do something with extra steps",
					Tags:        []string{"tag1", "tag2"},
					DueDate:     time.Now().Unix(),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "AddTask returns error",
			fields: fields{
				ddb: &ddbMock.MockDynamoDBClient{
					AddTaskErr: errors.New("test error"),
				},
				jwt: &tmMock.MockTokenManager{},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.AddTaskReq{
					Title:       "do something",
					Description: "do something with extra steps",
					Tags:        []string{"tag1", "tag2"},
					DueDate:     time.Now().Unix(),
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
			got, err := tr.AddTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todoServer.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateRecurringRule(t *testing.T) {
	type args struct {
		rule *proto.RecurringRule
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				rule: &proto.RecurringRule{
					CronExpression: "* * * * *",
					StartDate:      time.Now().Unix(),
					EndDate:        time.Now().AddDate(2, 0, 0).Unix(),
				},
			},
			wantErr: false,
		},
		{
			name: "invalid cron expression",
			args: args{
				rule: &proto.RecurringRule{
					CronExpression: "wrong, just wrong",
				},
			},
			wantErr: true,
		},
		{
			name: "another invalid cron expression",
			args: args{
				rule: &proto.RecurringRule{
					CronExpression: "0 100 0 0 0",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateRecurringRule(tt.args.rule); (err != nil) != tt.wantErr {
				t.Errorf("validateRecurringRule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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

func Test_todoServer_GetAllTasks(t *testing.T) {
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
			tr := &todoServer{
				UnimplementedTodoServer: tt.fields.UnimplementedTodoServer,
				ddb:                     tt.fields.ddb,
				jwt:                     tt.fields.jwt,
			}
			got, err := tr.GetAllTasks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
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
