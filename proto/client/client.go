package client

import (
	"context"
	"flag"
	"log"
	"time"

	pb "restapi/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "meen"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func StartGrpcClient() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.GetUserGrpc(ctx, &pb.UserRequest{})
	resp, err := c.GetUser(ctx, &pb.UserRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp)

}
