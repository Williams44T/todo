package interceptor

import (
	"context"
	"todo/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryAuthMiddleware authenticates JWTs.
// An "authorization" key set to the user's jwt must be provided in the metadata of the incoming context.
// Users attempting to sign up or sign in do not need a jwt in the metadata, but
// the "authorization" field should still exist in the metadata.
// A new jwt is issued and set in the header upon a successful call of the handler.
func (i *Interceptor) UnaryAuthMiddleware(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	// get jwt from metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata was not provided")
	}
	tokens := md.Get(common.AUTHORIZATION_METADATA_KEY)
	if len(tokens) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided in metadata")
	}

	switch info.FullMethod {
	case "Signup", "Signin":
		// TODO: Figure out these method names and implement alternative interceptor actions for each.
		// I might not implement incoming interceptor actions for these methods at all.
	default:
		// verify token and append user's ID to metadata
		userID, err := i.jwt.VerifyToken(tokens[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}
		md.Append(common.USERID_METADATA_KEY, userID)
		ctx = metadata.NewIncomingContext(ctx, md)
	}

	// call handler and return if error
	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	// get user id from metadata
	userIDs := md.Get(common.USERID_METADATA_KEY)
	if len(userIDs) == 0 {
		return nil, status.Error(codes.Unauthenticated, "user id is not provided in metadata")
	}

	// issue jwt and set it in the header
	jwt, err := i.jwt.IssueToken(userIDs[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to issue jwt: %v", err)
	}
	err = grpc.SetHeader(ctx, metadata.Pairs(common.JWT_METADATA_KEY, jwt))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to set jwt into header: %v", err)
	}

	return resp, nil
}
