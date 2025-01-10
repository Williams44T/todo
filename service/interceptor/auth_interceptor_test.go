package interceptor

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"todo/service/token_manager"
	"todo/service/token_manager/mock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// mockServerTransportStream allows grpc.SetHeader to function properly in tests
// https://stackoverflow.com/questions/74391247/not-able-to-set-header-when-calling-test-function-with-grpc-setheader-or-grpc-se
type mockServerTransportStream struct{}

func (m *mockServerTransportStream) Method() string                  { return "foo" }
func (m *mockServerTransportStream) SetHeader(md metadata.MD) error  { return nil }
func (m *mockServerTransportStream) SendHeader(md metadata.MD) error { return nil }
func (m *mockServerTransportStream) SetTrailer(md metadata.MD) error { return nil }

func TestInterceptor_UnaryAuthMiddleware(t *testing.T) {
	ctx := context.Background()
	ctx = grpc.NewContextWithServerTransportStream(ctx, &mockServerTransportStream{})
	validCtx := metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", "token"))

	type fields struct {
		jwt token_manager.TokenManagerInterface
	}
	type args struct {
		ctx     context.Context
		req     any
		info    *grpc.UnaryServerInfo
		handler grpc.UnaryHandler
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				jwt: &mock.MockTokenManager{
					TokenMap: map[string]string{"token": "user1234"},
				},
			},
			args: args{
				ctx:     validCtx,
				req:     nil,
				info:    &grpc.UnaryServerInfo{FullMethod: "SomeRPC"},
				handler: func(ctx context.Context, req any) (any, error) { return "response", nil },
			},
			want:    "response",
			wantErr: false,
		},
		{
			name: "VerifyToken returns error",
			fields: fields{jwt: &mock.MockTokenManager{
				VerifyTokenErr: errors.New("test error"),
			}},
			args: args{
				ctx:     validCtx,
				req:     nil,
				info:    &grpc.UnaryServerInfo{FullMethod: "SomeRPC"},
				handler: func(ctx context.Context, req any) (any, error) { return "response", nil },
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IssueToken returns error",
			fields: fields{jwt: &mock.MockTokenManager{
				IssueTokenErr: errors.New("test error"),
			}},
			args: args{
				ctx:     validCtx,
				req:     nil,
				info:    &grpc.UnaryServerInfo{FullMethod: "SomeRPC"},
				handler: func(ctx context.Context, req any) (any, error) { return "response", nil },
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interceptor{
				jwt: tt.fields.jwt,
			}
			got, err := i.UnaryAuthMiddleware(tt.args.ctx, tt.args.req, tt.args.info, tt.args.handler)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interceptor.UnaryAuthMiddleware() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interceptor.UnaryAuthMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
