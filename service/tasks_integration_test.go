//go:build integration
// +build integration

package service

import (
	"context"
	"testing"
	"time"
	"todo/dynamodb"
	proto "todo/proto/gen/service"
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
				ctx: context.Background(),
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
