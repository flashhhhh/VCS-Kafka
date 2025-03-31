package handler

import (
	"context"
	"user_service/internal/service"
	"user_service/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.EmptyResponse, error) {
	// Call the service to create a user
	err := h.service.CreateUser(ctx, req.Username, req.Password, req.Name)
	if err != nil {
		return nil, err
	}

	// Return an empty response
	return &pb.EmptyResponse{}, nil
}