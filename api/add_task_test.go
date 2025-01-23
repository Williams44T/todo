package api

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
	"todo/common"
	"todo/interfaces/dynamodb"
	ddbMock "todo/interfaces/dynamodb/mock"
	"todo/interfaces/token_manager"
	tmMock "todo/interfaces/token_manager/mock"
	proto "todo/proto/gen/go/api"

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
