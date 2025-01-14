package main

import (
	"os/exec"
	"testing"
	"time"
	"todo/cli/interceptor"

	"github.com/golang-jwt/jwt"
)

const todoCliPath = "../todo-cli"

func setAccessJWT(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "test-user",
		"exp": time.Now().Add(time.Duration(time.Minute * 5)).Unix(),
	})

	signed, err := token.SignedString([]byte("jwt_secret"))
	if err != nil {
		t.Errorf("failed to sign JWT: %v", err)
	}

	t.Setenv(interceptor.ACCESS_JWT_ENV_VAR, signed)
}

func Test_CLI_Integration(t *testing.T) {
	setAccessJWT(t)

	tests := []struct {
		name        string
		commandArgs []string
		wantErr     bool
	}{
		{
			name:        "happy path",
			commandArgs: []string{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(todoCliPath, tt.commandArgs...)
			output, err := cmd.CombinedOutput()
			if (err != nil) != tt.wantErr {
				t.Errorf("CLI Integration Test: error = %v, wantErr %v", err, tt.wantErr)
				t.Error("output:\n", string(output))
			}
		})
	}
}
