package interceptor

import (
	"fmt"
	"os"

	"todo/service/token_manager"
)

const (
	JWT_SECRET_ENV_KEY = "JWT_SECRET"
)

// Interceptor holds all the interceptor logic for the server
type Interceptor struct {
	jwt token_manager.TokenManagerInterface
}

// NewInterceptor returns a new instance of Interceptor
func NewInterceptor() (*Interceptor, error) {
	// get token manager
	jwtSecret, ok := os.LookupEnv(JWT_SECRET_ENV_KEY)
	if !ok {
		return nil, fmt.Errorf("%s must be provided as an environment variable", JWT_SECRET_ENV_KEY)
	}
	tokenManager, err := token_manager.NewTokenManager(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get token manager: %v", err)
	}

	return &Interceptor{
		jwt: tokenManager,
	}, nil
}
