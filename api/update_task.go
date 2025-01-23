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

func (t *todoServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskReq) (*proto.UpdateTaskResp, error) {
	// validate request
	if req.Task.Id == "" {
		return nil, errors.New("task id cannot be blank")
	}

	// get userid from ctx
	userIDs := metadata.ValueFromIncomingContext(ctx, common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("user id is not provided in metadata")
	}

	// update task
	var ddbRecurringRule *dynamodb.RecurringRule
	if req.Task.RecurringRule != nil {
		ddbRecurringRule = &dynamodb.RecurringRule{
			CronExpression: req.Task.RecurringRule.CronExpression,
			StartDate:      req.Task.RecurringRule.StartDate,
			EndDate:        req.Task.RecurringRule.EndDate,
		}
	}
	updateTaskResp, err := t.ddb.UpdateTask(ctx, &dynamodb.UpdateTaskReq{
		UserID: userIDs[0],
		TaskID: req.Task.Id,
		KVPairs: map[string]interface{}{
			dynamodb.TitleKey:         req.Task.Title,
			dynamodb.DescriptionKey:   req.Task.Title,
			dynamodb.StatusKey:        req.Task.Status.String(),
			dynamodb.TagsKey:          req.Task.Tags,
			dynamodb.ParentsKey:       req.Task.Parents,
			dynamodb.DueDateKey:       req.Task.DueDate,
			dynamodb.RecurringRuleKey: ddbRecurringRule,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %v", err)
	}

	// simplify
	ddbTask := updateTaskResp.Task

	// form and send response
	var recurringRule *proto.RecurringRule
	if ddbTask.RecurringRule != nil {
		recurringRule = &proto.RecurringRule{}
		recurringRule.CronExpression = ddbTask.RecurringRule.CronExpression
		recurringRule.StartDate = ddbTask.RecurringRule.StartDate
		recurringRule.EndDate = ddbTask.RecurringRule.EndDate
	}
	return &proto.UpdateTaskResp{
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
