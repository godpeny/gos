package delivery

import (
	"context"

	delivery_grpc "github.com/godpeny/gos/delivery/grpc"
)

// GetUser returns the User with request spec
func (s *GosServer) GetUser(ctx context.Context, req *delivery_grpc.UserRequest) (*delivery_grpc.UserResponse, error) {
	return &delivery_grpc.UserResponse{}, nil
}

// ListUser lists all users
func (s *GosServer) ListUser(ctx context.Context, req *delivery_grpc.UserRequest) (*delivery_grpc.UserListResponse, error) {
	return &delivery_grpc.UserListResponse{}, nil
}

// CreateUser
func (s *GosServer) CreateUser(ctx context.Context, req *delivery_grpc.UserRequest) (*delivery_grpc.UserResponse, error) {
	return &delivery_grpc.UserResponse{}, nil
}

// UpdateUser update User with request spec
func (s *GosServer) UpdateUser(ctx context.Context, req *delivery_grpc.UserRequest) (*delivery_grpc.UserResponse, error) {
	return &delivery_grpc.UserResponse{}, nil
}

// DeleteUser
func (s *GosServer) DeleteUser(ctx context.Context, req *delivery_grpc.UserRequest) (*delivery_grpc.UserResponse, error) {
	return &delivery_grpc.UserResponse{}, nil
}
