package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LogRequest(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("Incoming request. Method : %s. Request: %v", info.FullMethod, req)
	return handler(ctx, req)
}
