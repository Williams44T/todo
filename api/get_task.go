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

func (t *todoServer) GetTask(ctx context.Context, req *proto.GetTaskReq) (*proto.GetTaskResp, error) {
	// validate req
	if req.Id == "" {
		return nil, errors.New("id cannot be blank")
	}

	// get userid from ctx
	userIDs := metadata.ValueFromIncomingContext(ctx, common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("user id is not provided in metadata")
	}

	// get task
	getTaskResp, err := t.ddb.GetTask(ctx, &dynamodb.GetTaskReq{
		UserID: userIDs[0],
		TaskID: req.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	if getTaskResp.Task == nil {
		return nil, fmt.Errorf("task %s does not exist", req.Id)
	}

	// simplify
	ddbTask := getTaskResp.Task

	// form and send response
	var recurringRule *proto.RecurringRule
	if ddbTask.RecurringRule != nil {
		recurringRule = &proto.RecurringRule{}
		recurringRule.CronExpression = ddbTask.RecurringRule.CronExpression
		recurringRule.StartDate = ddbTask.RecurringRule.StartDate
		recurringRule.EndDate = ddbTask.RecurringRule.EndDate
	}
	return &proto.GetTaskResp{
		Task: &proto.Task{
			Id:            ddbTask.TaskID,
			Title:         ddbTask.Title,
			Description:   ddbTask.Description,
			Status:        proto.Status(proto.Status_value[ddbTask.Status]),
			Tags:          ddbTask.Tags,
			Parents:       ddbTask.Parents,
			DueDate:       ddbTask.DueDate,
			RecurringRule: recurringRule,
		},
	}, nil
}
