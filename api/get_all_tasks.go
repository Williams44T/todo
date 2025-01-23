package api

import (
	"context"
	"fmt"
	"todo/common"
	"todo/interfaces/dynamodb"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func (t *TodoServer) GetAllTasks(ctx context.Context, req *proto.GetAllTasksReq) (*proto.GetAllTasksResp, error) {
	// get userid from ctx
	userIDs := metadata.ValueFromIncomingContext(ctx, common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, fmt.Errorf("user id is not provided in metadata")
	}

	// get all tasks
	getAllTasksResp, err := t.ddb.GetAllTasks(ctx, &dynamodb.GetAllTasksReq{
		UserID: userIDs[0],
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks from ddb: %v", err)
	}
	tasks := []*proto.Task{}

	// convert all ddb tasks to proto tasks
	for _, task := range getAllTasksResp.Tasks {
		var recurringRule *proto.RecurringRule
		if task.RecurringRule != nil {
			recurringRule = &proto.RecurringRule{}
			recurringRule.CronExpression = task.RecurringRule.CronExpression
			recurringRule.StartDate = task.RecurringRule.StartDate
			recurringRule.EndDate = task.RecurringRule.EndDate
		}
		tasks = append(tasks, &proto.Task{
			Id:            task.TaskID,
			Title:         task.Title,
			Description:   task.Description,
			Status:        proto.Status(proto.Status_value[task.Status]),
			Tags:          task.Tags,
			Parents:       task.Parents,
			DueDate:       task.DueDate,
			RecurringRule: recurringRule,
		})
	}

	return &proto.GetAllTasksResp{
		Tasks: tasks,
	}, nil
}
