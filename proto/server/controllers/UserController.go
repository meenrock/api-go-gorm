package controllers

import (
	"context"
	"restapi/database"
	"restapi/models"

	pb "restapi/proto"
)

type GrpcServerImpl struct {
	pb.UnimplementedUserServiceServer
	pb.UnimplementedMealServiceServer
}

func (s *GrpcServerImpl) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	userResponses := make([]*pb.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &pb.UserResponse{
			FirstName: user.FirstName,
		}
	}
	response := &pb.UserResponse{}

	return response, nil
}
