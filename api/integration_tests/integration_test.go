//go:build integration

package integration_test

import (
	"context"
	"testing"
	"todo/api"
	"todo/common"
	proto "todo/proto/gen/go/api"

	"google.golang.org/grpc/metadata"
)

func Test_Integration_TodoServer(t *testing.T) {
	var todo *api.TodoServer
	var err error
	var userA string
	var userB string
	var taskA1 string
	var taskA2 string
	var taskB1 string
	var taskB2 string

	t.Run("Get Todo Server", func(t *testing.T) {
		todo, err = api.NewTodoServer(context.Background())
		if err != nil {
			t.Errorf("failed to get todo server: %v", err)
		}
	})

	t.Run("Signup UserA", func(t *testing.T) {
		resp, err := todo.Signup(context.Background(), &proto.SignupReq{
			FirstName: "userA",
			Email:     "userA@fake_email.com",
			Password:  "password",
		})
		if err != nil {
			t.Errorf("failed to sign up: %v", err)
		}
		if resp.AccessJWT == "" {
			t.Error("no access jwt returned in signup resp")
		}
		if resp.UserID == "" {
			t.Error("no user id returned in signup resp")
		}
		userA = resp.UserID
	})

	t.Run("Signup UserB", func(t *testing.T) {
		resp, err := todo.Signup(context.Background(), &proto.SignupReq{
			FirstName: "userB",
			Email:     "userB@fake_email.com",
			Password:  "password",
		})
		if err != nil {
			t.Errorf("failed to sign up: %v", err)
		}
		if resp.AccessJWT == "" {
			t.Error("no access jwt returned in signup resp")
		}
		if resp.UserID == "" {
			t.Error("no user id returned in signup resp")
		}
		userB = resp.UserID
	})

	t.Run("UserA adds tasks", func(t *testing.T) {
		// Verifying the user's JWT is not tested throughout this test since none of the interceptors are involved,
		// so we mock the action of the interceptor placing the user's id in the metadata after successful auth.
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userA))
		resp, err := todo.AddTask(ctx, &proto.AddTaskReq{Title: "taskA1"})
		if err != nil {
			t.Errorf("failed to add task 1: %v", err)
		}
		if resp.Id == "" {
			t.Error("no task id returned in AddTask resp")
		}
		taskA1 = resp.Id
		resp, err = todo.AddTask(ctx, &proto.AddTaskReq{Title: "taskA2"})
		if err != nil {
			t.Errorf("failed to add task 2: %v", err)
		}
		if resp.Id == "" {
			t.Error("no task id returned in AddTask resp")
		}
		taskA2 = resp.Id
	})

	t.Run("UserB adds tasks", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userB))
		resp, err := todo.AddTask(ctx, &proto.AddTaskReq{Title: "taskB1"})
		if err != nil {
			t.Errorf("failed to add task 1: %v", err)
		}
		if resp.Id == "" {
			t.Error("no task id returned in AddTask resp")
		}
		taskB1 = resp.Id
		resp, err = todo.AddTask(ctx, &proto.AddTaskReq{Title: "taskB2"})
		if err != nil {
			t.Errorf("failed to add task 2: %v", err)
		}
		if resp.Id == "" {
			t.Error("no task id returned in AddTask resp")
		}
		taskB2 = resp.Id
	})

	t.Run("UserA updates a task", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userA))
		getTaskResp, err := todo.GetTask(ctx, &proto.GetTaskReq{Id: taskA1})
		if err != nil {
			t.Errorf("failed to GetTask: %v:", err)
		}

		updatedTask := getTaskResp.Task
		updatedTask.Tags = []string{"tag1"}

		updateTaskResp, err := todo.UpdateTask(ctx, &proto.UpdateTaskReq{Task: updatedTask})
		if err != nil {
			t.Errorf("failed to update task: %v", err)
		}
		if updateTaskResp.Task.Tags[0] != updatedTask.Tags[0] {
			t.Error("unexpected returned Task from UpdateTask")
		}
	})

	t.Run("UserA attempts to update UserB's task", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userA))
		resp, _ := todo.UpdateTask(ctx, &proto.UpdateTaskReq{
			Task: &proto.Task{Id: taskB1, Title: "I've updated your task, ha ha!"},
		})
		if resp != nil {
			t.Errorf("unauthorized UpdateTask call was successful, resp: %v:", resp)
		}
	})

	t.Run("Confirm UserB's task was not updated by UserA", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userB))
		resp, err := todo.GetTask(ctx, &proto.GetTaskReq{Id: taskB1})
		if err != nil {
			t.Errorf("failed to GetTask: %v:", err)
		}
		if resp.Task.Title != "taskB1" {
			t.Error("unauthorized UpdateTask call was successful")
		}
	})

	t.Run("UserA deletes a task", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userA))
		_, err := todo.DeleteTask(ctx, &proto.DeleteTaskReq{TaskId: taskA1})
		if err != nil {
			t.Errorf("failed to delete task: %v", err)
		}

		resp, err := todo.GetAllTasks(ctx, &proto.GetAllTasksReq{})
		if err != nil {
			t.Errorf("failed to GetAllTasks: %v", err)
		}
		if resp.Tasks[0].Id != taskA2 || len(resp.Tasks) > 1 {
			t.Errorf("taskA1 was not deleted successfully")
		}
	})

	t.Run("UserA attempts to delete UserB's task", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userA))
		_, err := todo.DeleteTask(ctx, &proto.DeleteTaskReq{TaskId: taskB1})
		if err != nil {
			// the behaviour here should be rethought
			t.Errorf("DeleteTask shouldn't error here: %v", err)
		}
	})

	t.Run("Confirm UserB's task was not updated", func(t *testing.T) {
		ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(common.USERID_METADATA_KEY, userB))
		resp, err := todo.GetAllTasks(ctx, &proto.GetAllTasksReq{})
		if err != nil {
			t.Errorf("failed to GetAllTasks: %v:", err)
		}
		userBTaskIdSet := map[string]struct{}{
			taskB1: {},
			taskB2: {},
		}
		gotTaskIdSet := map[string]struct{}{}
		for _, task := range resp.Tasks {
			gotTaskIdSet[task.Id] = struct{}{}
		}
		for id := range userBTaskIdSet {
			if _, ok := gotTaskIdSet[id]; !ok {
				t.Errorf("task %s was not returned from GetAllTasks", id)
			}
		}
	})

}
