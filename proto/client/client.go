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

func GetUserById(client pb.UserServiceClient) {

}

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
	resp, err := c.CreateUser(ctx, &pb.UserCreateRequest{
		FirstName: "HellUpdate",
		LastName:  "MeUpdate",
		Weight:    67.00,
		Height:    170.00,
	})
	user, err2 := c.GetUserById(ctx, &pb.UserRequest{
		Id: 4,
	})
	update, err3 := c.EditUser(ctx, &pb.UserEditRequest{
		Id:        5,
		FirstName: "HellUpdate",
		LastName:  "MeUpdate",
	})
	delete, err4 := c.DeleteUser(ctx, &pb.UserRequest{
		Id: 1,
	})

	log.Printf("CreateUser: %s", resp, err)
	log.Printf("UserList: %s", user, err2)
	log.Printf("Update: %s", update, err3)
	log.Printf("Delete: %s", delete, err4)

}
