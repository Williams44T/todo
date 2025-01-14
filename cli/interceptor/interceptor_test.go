package interceptor

import (
	"context"
	"testing"

	"google.golang.org/grpc"
)

func TestInterceptor_UnaryAuthMiddleware(t *testing.T) {
	type args struct {
		ctx     context.Context
		method  string
		req     interface{}
		reply   interface{}
		cc      *grpc.ClientConn
		invoker grpc.UnaryInvoker
		opts    []grpc.CallOption
	}
	tests := []struct {
		name    string
		i       *Interceptor
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				ctx:     context.Background(),
				method:  "AnyRPC",
				req:     nil,
				reply:   nil,
				cc:      &grpc.ClientConn{},
				invoker: func(context.Context, string, any, any, *grpc.ClientConn, ...grpc.CallOption) error { return nil },
				opts:    []grpc.CallOption{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interceptor{}
			if err := i.UnaryAuthMiddleware(tt.args.ctx, tt.args.method, tt.args.req, tt.args.reply, tt.args.cc, tt.args.invoker, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Interceptor.UnaryAuthMiddleware() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
