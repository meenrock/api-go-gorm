package proto

import (
	"context"
	"fmt"

	pb "restapi"

	"google.golang.org/grpc"
)

type GrpcServerImpl struct {
}

func (s *GrpcServerImpl) StartPing(ctx context.Context) error {
	fmt.Println("Ping Received")

	return nil
}

func StartPingPongServer() {
	grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()

	// Create a gRPC client
	grpcClient := pb.NewMyServiceClient(grpcConn)
}
