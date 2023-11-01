package grpcserver

import (
	// Import required packages
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	userpb "main.go/userProto"
)

// Mock database
var mockDB = []userpb.User{
	{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{Id: 2, Fname: "Rajendro", City: "Kolkata", Phone: 8250771252, Height: 5.9, Married: false},

	// Add more users if needed
}

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUserById(ctx context.Context, req *userpb.UserRequest) (*userpb.User, error) {
	id := req.GetId()
	for _, user := range mockDB {
		if user.Id == int32(id) {
			return &user, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("User with ID %d not found", id))
}

func (s *server) GetUsersByIds(ctx context.Context, req *userpb.UsersRequest) (*userpb.Users, error) {
	ids := req.GetIds()
	var users []*userpb.User
	for _, id := range ids {
		for _, user := range mockDB {
			if user.Id == id {
				users = append(users, &user)
				break
			}
		}
	}
	return &userpb.Users{Users: users}, nil
}
func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
