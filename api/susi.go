package api

import (
	"context"
	"errors"
	"fmt"
	"todo/interfaces/dynamodb"
	proto "todo/proto/gen/go/api"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
)

// hashPassword uses argon2id default params to salt and hash a user's given password.
// https://github.com/alexedwards/argon2id
func hashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Signup hashes the password, generates a user id, and then adds the user to the database,
// before returning an access jwt and the user id.
func (t *TodoServer) Signup(ctx context.Context, req *proto.SignupReq) (*proto.SignupResp, error) {
	// validate request
	if req.FirstName == "" {
		return nil, errors.New("first name cannot be blank")
	}
	// TODO: implement stronger email validation
	if req.Email == "" {
		return nil, errors.New("email cannot be blank")
	}
	// TODO: implement stronger password validation
	if req.Password == "" {
		return nil, errors.New("password cannot be blank")
	}

	// hash password
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// generate user id
	userID := uuid.New()

	// add user to DDB
	_, err = t.ddb.AddUser(ctx, &dynamodb.AddUserReq{
		User: dynamodb.User{
			ID:             hashedPassword,
			FirstName:      req.FirstName,
			LastName:       req.LastName,
			Email:          req.Email,
			HashedPassword: hashedPassword,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add user: %v", err)
	}

	// generate access token
	token, err := t.jwt.IssueToken(userID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to issue jwt: %v", err)
	}

	// TODO: generate refresh token

	return &proto.SignupResp{
		AccessJWT: token,
		UserID:    userID.String(),
	}, nil
}

// Signin gets the user with the given id, compares the given password and the stored hashed password,
// and returns an access jwt if there is a match.
func (t *TodoServer) Signin(ctx context.Context, req *proto.SigninReq) (*proto.SigninResp, error) {
	// validate request
	if req.UserID == "" {
		return nil, errors.New("userid cannot be blank")
	}

	// get user's stored hash password
	getUserResp, err := t.ddb.GetUser(ctx, &dynamodb.GetUserReq{
		ID: req.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	if getUserResp.User == nil {
		return nil, fmt.Errorf("user %s does not exist", req.UserID)
	}

	// compare password
	match, err := argon2id.ComparePasswordAndHash(req.Password, getUserResp.User.HashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to compare password to hash: %v", err)
	}
	if !match {
		return nil, errors.New("invalid password")
	}

	// generate access token
	token, err := t.jwt.IssueToken(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to issue jwt: %v", err)
	}

	// TODO: generate refresh token

	return &proto.SigninResp{
		AccessJWT: token,
	}, nil
}
