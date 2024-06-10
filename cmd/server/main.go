package main

import (
	"context"
	"log"
	"net"

	pb "github.com/uulwake/grpc/generated/grpc"
	pbCommon "github.com/uulwake/grpc/generated/grpc/common"
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

type order struct {
	pbCommon.UnimplementedOrderServer
}

func (o *order) GetUserOrders(ctx context.Context, req *pbCommon.UserId) (*pbCommon.Orders, error) {
	return &pbCommon.Orders{
		Orders: []*pbCommon.OrderData{
			{
				ID: 1,
				RecipientName: "ulrich",
				Address: 12.4,
				Status: pbCommon.ORDER_STATUS_CANCELLED,
				RecipientPhone: "1234",
			},
		},
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
	pbCommon.RegisterOrderServer(s, &order{})
	log.Printf("server is listening at port %v", listener.Addr())

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
