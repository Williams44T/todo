package api

import (
	"context"
	"errors"
	"fmt"

	"todo/common"
	"todo/interfaces/dynamodb"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func (t *TodoServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskReq) (*proto.DeleteTaskResp, error) {
	// validate request
	if req.TaskId == "" {
		return nil, errors.New("task id cannot be blank")
	}

	// get userid from ctx
	userIDs := metadata.ValueFromIncomingContext(ctx, common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("user id is not provided in metadata")
	}

	_, err := t.ddb.DeleteTask(ctx, &dynamodb.DeleteTaskReq{
		UserID: userIDs[0],
		TaskID: req.TaskId,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to delete task: %v", err)
	}

	return &proto.DeleteTaskResp{}, nil
}
