package api

import (
	"context"
	"errors"
	"fmt"
	"todo/common"
	"todo/interfaces/dynamodb"
	proto "todo/proto/gen/go/api"

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

func (t *TodoServer) AddTask(ctx context.Context, req *proto.AddTaskReq) (*proto.AddTaskResp, error) {
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

	// generate task id
	taskID := uuid.New().String()

	// use db client to add task
	ddbRecurringRule := &dynamodb.RecurringRule{}
	if req.RecurringRule != nil {
		ddbRecurringRule.CronExpression = req.RecurringRule.CronExpression
		ddbRecurringRule.StartDate = req.RecurringRule.StartDate
		ddbRecurringRule.EndDate = req.RecurringRule.EndDate
	}
	_, err := t.ddb.AddTask(ctx, &dynamodb.AddTaskReq{
		Task: dynamodb.Task{
			UserID:        userIDs[0],
			TaskID:        taskID,
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

	return &proto.AddTaskResp{
		Id: taskID,
	}, nil
}
