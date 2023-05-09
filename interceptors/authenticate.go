package interceptors

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Authenticate(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if info.FullMethod == "/helloworld.Greeter/SayHello" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(401, "token is missing")
	}

	token := md.Get("token")[0]

	if token != "valid" {
		return nil, status.Error(401, "token is not valid")
	}

	return handler(ctx, req)
}
