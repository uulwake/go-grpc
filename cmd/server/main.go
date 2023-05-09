package main

import (
	"context"
	"log"
	"net"

	pb "github.com/uulwake/grpc/generated"
	"github.com/uulwake/grpc/interceptors"
	"google.golang.org/grpc"
)

type greeter struct {
	pb.UnimplementedGreeterServer
}

func (g *greeter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

type user struct {
	pb.UnimplementedUserServer
}

func (u *user) GetUserByID(ctx context.Context, req *pb.ID) (*pb.UserData, error) {
	return &pb.UserData{
		ID:    int64(1),
		Name:  "Ulrich",
		Email: "ulrich@gmail.com",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.LogRequest,
			interceptors.Authenticate,
		),
	)
	pb.RegisterGreeterServer(s, &greeter{})
	pb.RegisterUserServer(s, &user{})
	log.Printf("server is listening at port %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
