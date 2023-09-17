package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"restapi/database"
	"restapi/models"
	pb "restapi/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type GrpcServerImpl struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedMealServiceServer
}

func (s *GrpcServerImpl) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	if err := db.First(&user, req.Id).Error; err != nil {
		return nil, err
	}

	// userResponses := make([]*pb.UserResponse, len(users))
	// for i, user := range users {
	// 	userResponses[i] = &pb.UserResponse{
	// 		FirstName: user.FirstName,
	// 		LastName:  user.LastName,
	// 	}
	// }
	// response := &pb.UserResponse{}

	response := &pb.UserResponse{
		Id:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return response, nil
}

func (s *GrpcServerImpl) CreateUser(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserResponse, error) {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	db.Create(&user)

	// userResponses := make([]*pb.UserResponse, len(users))
	// for i, user := range users {
	// 	userResponses[i] = &pb.UserResponse{
	// 		FirstName: user.FirstName,
	// 		LastName:  user.LastName,
	// 	}
	// }
	response := &pb.UserResponse{
		Id:        int32(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
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
