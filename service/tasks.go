package service

import (
	"context"
	"errors"
	"fmt"

	"todo/common"
	"todo/dynamodb"
	proto "todo/proto/gen/service"

	"github.com/adhocore/gronx"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

func validateRecurringRule(rule *proto.RecurringRule) error {
	if rule == nil {
		return nil
	}
	if !gronx.IsValid(rule.CronExpression) {
		return errors.New("invalid cron expression")
	}
	return nil
}

func (t *todoServer) AddTask(ctx context.Context, req *proto.AddTaskReq) (*proto.AddTaskResp, error) {
	// validate req
	if req.Title == "" {
		return nil, errors.New("title cannot be blank")
	}
	if err := validateRecurringRule(req.RecurringRule); err != nil {
		return nil, fmt.Errorf("invalid recurring rule: %v", err)
	}

	// get userid from ctx
	userIDs := metadata.ValueFromIncomingContext(ctx, common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("user id is not provided in metadata")
	}

	// use db client to add task
	ddbRecurringRule := &dynamodb.RecurringRule{}
	if req.RecurringRule != nil {
		ddbRecurringRule.CronExpression = req.RecurringRule.CronExpression
		ddbRecurringRule.StartDate = req.RecurringRule.StartDate
		ddbRecurringRule.EndDate = req.RecurringRule.EndDate
	}
	_, err := t.ddb.AddTask(ctx, &dynamodb.AddTaskReq{
		Task: dynamodb.Task{
			ID:            uuid.New().String(),
			UserID:        userIDs[0],
			Title:         req.Title,
			Description:   req.Description,
			Status:        req.Status.String(),
			Tags:          req.Tags,
			Parents:       req.Parents,
			RecurringRule: ddbRecurringRule,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add task to database: %v", err)
	}

	return &proto.AddTaskResp{}, nil
}

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
		ID: req.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %v", err)
	}
	if getTaskResp.Task == nil {
		return nil, fmt.Errorf("task %s does not exist", req.Id)
	}

	// simplify
	ddbTask := getTaskResp.Task

	// compare user ids
	if ddbTask.UserID != userIDs[0] {
		return nil, errors.New("user is not authorized to read this task")
	}

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
			Id:            ddbTask.ID,
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
