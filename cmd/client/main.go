package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/uulwake/grpc/generated"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	md := metadata.New(map[string]string{"token": "valid"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.GetUserByID(ctx, &pb.ID{ID: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(resp.ID, resp.Name, resp.Email)

}
