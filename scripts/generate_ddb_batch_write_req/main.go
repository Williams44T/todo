/*
This script is used by github workflows to populate the Local DynamoDB with entries
necessary for integration tests to pass.
*/
package main

import (
	"encoding/json"
	"log"
	"os"
	"todo/common"

	"github.com/alexedwards/argon2id"
)

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html
type ID struct {
	S string `json:"S,omitempty"`
}
type HashedPassword struct {
	S string `json:"S,omitempty"`
}
type UserID struct {
	S string `json:"S,omitempty"`
}
type TaskID struct {
	S string `json:"S,omitempty"`
}
type Status struct {
	S string `json:"S,omitempty"`
}
type Item struct {
	// user fields
	ID             *ID             `json:"id,omitempty"`
	HashedPassword *HashedPassword `json:"hashed_password,omitempty"`

	// task fields
	UserID *UserID `json:"user_id,omitempty"`
	TaskID *TaskID `json:"task_id,omitempty"`
	Status *Status `json:"status,omitempty"`
}
type PutRequest struct {
	Item *Item `json:"Item,omitempty"`
}
type WriteRequest struct {
	PutRequest *PutRequest `json:"PutRequest,omitempty"`
}
type RequestItems struct {
	UsersTable []*WriteRequest `json:"todo-users,omitempty"`
	TasksTable []*WriteRequest `json:"todo-tasks,omitempty"`
}

func getPutUserRequest(id, hashedPassword string) *PutRequest {
	return &PutRequest{
		Item: &Item{
			ID: &ID{
				S: id,
			},
			HashedPassword: &HashedPassword{
				S: hashedPassword,
			},
		},
	}
}

func getPutTaskRequest(userID, taskID, status string) *PutRequest {
	return &PutRequest{
		Item: &Item{
			UserID: &UserID{
				S: userID,
			},
			TaskID: &TaskID{
				S: taskID,
			},
			Status: &Status{
				S: status,
			},
		},
	}
}

// hashPassword uses argon2id default params to salt and hash a user's given password.
// https://github.com/alexedwards/argon2id
func hashPassword(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatalf("failed to hash password: %v", err)
	}
	return hash
}

func main() {
	data := RequestItems{
		UsersTable: []*WriteRequest{
			{PutRequest: getPutUserRequest(common.TEST_USER_1_ID, hashPassword(common.TEST_USER_1_PASSWORD))},
		},
		TasksTable: []*WriteRequest{
			{PutRequest: getPutTaskRequest(common.TEST_USER_1_ID, common.TASK_1A_ID, "INCOMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_1_ID, common.TASK_1B_ID, "INCOMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_1_ID, common.TASK_1C_ID, "COMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_1_ID, common.TASK_1D_ID, "COMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_2_ID, common.TASK_2A_ID, "INCOMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_2_ID, common.TASK_2B_ID, "INCOMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_2_ID, common.TASK_2C_ID, "COMPLETE")},
			{PutRequest: getPutTaskRequest(common.TEST_USER_2_ID, common.TASK_2D_ID, "COMPLETE")},
		},
	}

	file, err := os.Create("interfaces/dynamodb/mock/batch_write_req.json")
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: Pretty-print JSON

	err = encoder.Encode(data)
	if err != nil {
		log.Fatalf("failed to encode data: %v", err)
	}
}
