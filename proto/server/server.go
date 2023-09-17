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

func (s *GrpcServerImpl) GetUser(req *pb.Empty, stream pb.UserService_GetUserServer) error {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil
	}
	defer db.Close()

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil
	}

	// userResponses := make([]*pb.UserResponseAll, len(users))
	// for i, user := range users {
	// 	userResponses[i] = &pb.UserResponseAll{
	// 		Id:           int32(user.ID),
	// 		FirstName:    user.FirstName,
	// 		LastName:     user.LastName,
	// 		NickName:     user.NickName,
	// 		PhoneNumber:  user.PhoneNumber,
	// 		Email:        user.Email,
	// 		Sex:          user.Sex,
	// 		Age:          int32(user.Age),
	// 		Height:       float32(user.Height),
	// 		Weight:       float32(user.Weight),
	// 		AllergicFood: user.AllergicFood,
	// 		FavFood:      user.FavFood,
	// 		ExpectedBMI:  float32(user.ExpectedBMI),
	// 	}
	// }

	for _, user := range users {
		userResponses := &pb.UserResponseAll{
			Id:           int32(user.ID),
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			NickName:     user.NickName,
			PhoneNumber:  user.PhoneNumber,
			Email:        user.Email,
			Sex:          user.Sex,
			Age:          int32(user.Age),
			Height:       float32(user.Height),
			Weight:       float32(user.Weight),
			AllergicFood: user.AllergicFood,
			FavFood:      user.FavFood,
			ExpectedBMI:  float32(user.ExpectedBMI),
		}

		if err := stream.Send(userResponses); err != nil {
			return err
		}
	}

	return nil
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
		Id:           int32(user.ID),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		NickName:     user.NickName,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		Sex:          user.Sex,
		Age:          int32(user.Age),
		Height:       float32(user.Height),
		Weight:       float32(user.Weight),
		AllergicFood: user.AllergicFood,
		FavFood:      user.FavFood,
		ExpectedBMI:  float32(user.ExpectedBMI),
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
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		NickName:     req.NickName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Sex:          req.Sex,
		Age:          int(req.Age),
		Height:       float32(req.Height),
		Weight:       float32(req.Weight),
		AllergicFood: req.AllergicFood,
		FavFood:      req.FavFood,
		ExpectedBMI:  float32(req.ExpectedBMI),
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	db.Create(&user)

	response := &pb.UserResponse{
		Id:           int32(user.ID),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		NickName:     user.NickName,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		Sex:          user.Sex,
		Age:          int32(user.Age),
		Height:       float32(user.Height),
		Weight:       float32(user.Weight),
		AllergicFood: user.AllergicFood,
		FavFood:      user.FavFood,
		ExpectedBMI:  float32(user.ExpectedBMI),
	}

	return response, nil
}

func (s *GrpcServerImpl) EditUser(ctx context.Context, req *pb.UserEditRequest) (*pb.UserResponse, error) {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &models.User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		NickName:     req.NickName,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		Sex:          req.Sex,
		Age:          int(req.Age),
		Height:       float32(req.Height),
		Weight:       float32(req.Weight),
		AllergicFood: req.AllergicFood,
		FavFood:      req.FavFood,
		ExpectedBMI:  float32(req.ExpectedBMI),
	}

	if err := db.First(user, req.Id).Error; err != nil {
		return nil, err
	}

	db.Save(&user)

	response := &pb.UserResponse{
		Id:           int32(user.ID),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		NickName:     user.NickName,
		PhoneNumber:  user.PhoneNumber,
		Email:        user.Email,
		Sex:          user.Sex,
		Age:          int32(user.Age),
		Height:       float32(user.Height),
		Weight:       float32(user.Weight),
		AllergicFood: user.AllergicFood,
		FavFood:      user.FavFood,
		ExpectedBMI:  float32(user.ExpectedBMI),
	}

	return response, nil
}

func (s *GrpcServerImpl) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.UserDeleteConfirmation, error) {

	db, err := database.ConnectDBPostgres()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.User
	if err := db.First(&user, req.Id).Error; err != nil {
		return nil, err
	}

	db.Delete(&user)

	response := &pb.UserDeleteConfirmation{
		Id: int32(user.ID),
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
