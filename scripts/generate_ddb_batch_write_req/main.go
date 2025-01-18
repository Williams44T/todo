/*
This script is used by github workflows to populate the Local DynamoDB with entries
necessary for integration tests to pass.
*/
package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/alexedwards/argon2id"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BatchWriteItem.html
type ID struct {
	S string `json:"S"`
}
type HashedPassword struct {
	S string `json:"S"`
}
type Item struct {
	ID             ID             `json:"ID"`
	HashedPassword HashedPassword `json:"HashedPassword"`
}
type PutRequest struct {
	Item Item `json:"Item"`
}
type WriteRequest struct {
	PutRequest PutRequest `json:"PutRequest"`
}
type RequestItems struct {
	Todo_users []WriteRequest `json:"Todo_users"`
}

func getPutRequest(id string, hashedPassword string) PutRequest {
	return PutRequest{
		Item: Item{
			ID: ID{
				S: id,
			},
			HashedPassword: HashedPassword{
				S: hashedPassword,
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
		Todo_users: []WriteRequest{
			{PutRequest: getPutRequest("test-user-1", hashPassword("password"))},
		},
	}

	file, err := os.Create("dynamodb/mock/batch_write_req.json")
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
