package interceptor

import (
	"fmt"
	"os"

	"todo/common"
	"todo/service/token_manager"
)

// Interceptor holds all the interceptor logic for the server
type Interceptor struct {
	jwt token_manager.TokenManagerInterface
}

// NewInterceptor returns a new instance of Interceptor
func NewInterceptor() (*Interceptor, error) {
	// get token manager
	jwtSecret, ok := os.LookupEnv(common.JWT_SECRET_ENV_VAR)
	if !ok {
		return nil, fmt.Errorf("%s must be provided as an environment variable", common.JWT_SECRET_ENV_VAR)
	}
	tokenManager, err := token_manager.NewTokenManager(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get token manager: %v", err)
	}

	return &Interceptor{
		jwt: tokenManager,
	}, nil
}
