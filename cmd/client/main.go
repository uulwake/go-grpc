package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/uulwake/grpc/generated/grpc"
)

func getUserByID(conn *grpc.ClientConn) error {
	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.New(map[string]string{"token": "valid"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.GetUserByID(ctx, &pb.ID{ID: 1})
	if err != nil {
		return err
	}

	log.Println("GetUserByID: ", resp.ID, resp.Name, resp.Email)
	return nil
}

func sayHello(conn *grpc.ClientConn) error {
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Ulrich"})
	if err != nil {
		return err
	}

	log.Println("SayHello: ", resp.Message)
	return nil
}

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	err = getUserByID(conn)
	if err != nil {
		log.Fatalf("cannot getUserByID: %v", err)
	}

	err = sayHello(conn)
	if err != nil {
		log.Fatalf("cannot sayHello: %v", err)
	}
}
