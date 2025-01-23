//go:build integration
// +build integration

package api

import (
	"context"
	"testing"
	"time"
	"todo/common"
	"todo/interfaces/dynamodb"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func Test_Integration_todoServer_AddTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.AddTaskReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.AddTaskReq{
					Title:       "do something",
					Description: "do something with extra steps",
					Tags:        []string{"tag1", "tag2"},
					DueDate:     time.Now().Unix(),
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
				t.Errorf("integration todoServer.AddTask() failed to get database client: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
			}

			_, err = tr.AddTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Integration_todoServer_GetTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetTaskReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.GetTaskReq{
					Id: common.TASK_1A_ID,
				},
			},
			wantErr: false,
		},
		{
			name: "task does not exist",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.GetTaskReq{
					Id: "fake_task_id",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.GetTask() failed to get database client: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
			}

			_, err = tr.GetTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Integration_todoServer_GetAllTasks(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.GetAllTasksReq
	}
	tests := []struct {
		name          string
		args          args
		wantTaskIdSet map[string]struct{}
		wantErr       bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				req: &proto.GetAllTasksReq{},
			},
			wantTaskIdSet: map[string]struct{}{
				common.TASK_2A_ID: {},
				common.TASK_2B_ID: {},
				common.TASK_2C_ID: {},
				common.TASK_2D_ID: {},
			},
			wantErr: false,
		},
		{
			name: "no tasks tied to provided user id",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, "fake_user_id")),
				req: &proto.GetAllTasksReq{},
			},
			wantTaskIdSet: map[string]struct{}{},
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.GetAllTasks() failed to get database client: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
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

func Test_Integration_todoServer_UpdateTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.UpdateTaskReq
	}
	tests := []struct {
		name    string
		args    args
		want    *proto.UpdateTaskResp
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{
						Userid: common.TEST_USER_1_ID,
						Id:     common.TASK_1A_ID,
						Status: proto.Status_COMPLETE,
					},
				},
			},
			want: &proto.UpdateTaskResp{
				Task: &proto.Task{
					Userid: common.TEST_USER_1_ID,
					Id:     common.TASK_1A_ID,
					Status: proto.Status_COMPLETE,
				},
			},
			wantErr: false,
		},
		{
			name: "task does not exist",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.UpdateTaskReq{
					Task: &proto.Task{
						Userid: common.TEST_USER_1_ID,
						Id:     "nonexistent task id",
						Status: proto.Status_COMPLETE,
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.UpdateTask() failed to get database client: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
			}

			_, err = tr.UpdateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_Integration_todoServer_DeleteTask(t *testing.T) {
	type args struct {
		ctx context.Context
		req *proto.DeleteTaskReq
	}
	tests := []struct {
		name          string
		args          args
		want          *proto.DeleteTaskResp
		wantTaskIdSet map[string]struct{}
		wantErr       bool
	}{
		{
			name: "happy path",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				req: &proto.DeleteTaskReq{
					TaskId: common.TASK_2A_ID,
				},
			},
			want: &proto.DeleteTaskResp{},
			wantTaskIdSet: map[string]struct{}{
				// common.TASK_2A_ID: {}, #deleted
				common.TASK_2B_ID: {},
				common.TASK_2C_ID: {},
				common.TASK_2D_ID: {},
			},
			wantErr: false,
		},
		{
			name: "happy path - nonexistent task doesn't throw error",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				req: &proto.DeleteTaskReq{
					TaskId: "nonexistent_id",
				},
			},
			want: &proto.DeleteTaskResp{},
			wantTaskIdSet: map[string]struct{}{
				// common.TASK_2A_ID: {}, # deleted in the first DeleteTask test
				common.TASK_2B_ID: {},
				common.TASK_2C_ID: {},
				common.TASK_2D_ID: {},
			},
			wantErr: false,
		},
		{
			name: "happy path - can't delete a task you don't own",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_1_ID)),
				req: &proto.DeleteTaskReq{
					TaskId: common.TASK_2A_ID,
				},
			},
			want: &proto.DeleteTaskResp{},
			wantTaskIdSet: map[string]struct{}{
				// common.TASK_2A_ID: {}, # deleted in the first DeleteTask test
				common.TASK_2B_ID: {},
				common.TASK_2C_ID: {},
				common.TASK_2D_ID: {},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get database client
			databaseClient, err := dynamodb.NewDynamoDBClient(tt.args.ctx)
			if err != nil {
				t.Errorf("integration todoServer.DeleteTask() failed to get database client: %v", err)
			}

			tr := &todoServer{
				ddb: databaseClient,
			}

			_, err = tr.DeleteTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("todoServer.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// make sure task was actually deleted
			resp, err := tr.GetAllTasks(
				metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, common.TEST_USER_2_ID)),
				&proto.GetAllTasksReq{},
			)
			if err != nil {
				t.Errorf("todoServer.GetAllTasks() error = %v", err)
			}
			gotTaskIdSet := map[string]struct{}{}
			for _, gotTask := range resp.Tasks {
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
