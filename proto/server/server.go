package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "restapi/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServerImpl struct {
	pb.UnimplementedUserServiceServer
}

func (s *GrpcServerImpl) GetUserGrpc(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	response := &pb.UserResponse{
		FirstName: "Supawat",
	}

	return response, nil
}

func StartGrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}

	// Create a gRPC client
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &GrpcServerImpl{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
