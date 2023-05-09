package interceptors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Authenticate(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if info.FullMethod == "/helloworld.Greeter/SayHello" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata when authenticate")
	}

	token := md.Get("token")[0]

	if token != "valid" {
		return nil, errors.New("rpc is not authenticate")
	}

	return handler(ctx, req)
}
