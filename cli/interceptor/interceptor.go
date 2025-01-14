package interceptor

import (
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	ACCESS_JWT_ENV_VAR = "TODO_SERVICE_ACCESS_JWT"
)

// Interceptor holds all the interceptor logic for the CLI
type Interceptor struct{}

// NewInterceptor returns a new instance of Interceptor
func NewInterceptor() (*Interceptor, error) {
	return &Interceptor{}, nil
}

// UnaryAuthMiddleware adds the existing access jwt to the outgoing metadata.
func (i *Interceptor) UnaryAuthMiddleware(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", os.Getenv(ACCESS_JWT_ENV_VAR))
	return invoker(ctx, method, req, reply, cc, opts...)
}
