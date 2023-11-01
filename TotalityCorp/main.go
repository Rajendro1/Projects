package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
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

func startGRPCServer() {
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

func startGinServer() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		// idStr := c.Param("id")
		idStr := c.Request.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Connect to gRPC server
		cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
			return
		}
		defer cc.Close()

		client := userpb.NewUserServiceClient(cc)
		resp, err := client.GetUserById(context.Background(), &userpb.UserRequest{Id: int32(id)})
		if err != nil {
			if status.Code(err) == codes.NotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user details"})
			return
		}

		c.JSON(http.StatusOK, resp)
	})
	r.GET("/users", func(c *gin.Context) {
		idsStr := c.Query("ids")
		idStrs := strings.Split(idsStr, ",")
		var ids []int32
		for _, idStr := range idStrs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
				return
			}
			ids = append(ids, int32(id))
		}

		// Connect to gRPC server
		cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
			return
		}
		defer cc.Close()

		client := userpb.NewUserServiceClient(cc)
		resp, err := client.GetUsersByIds(context.Background(), &userpb.UsersRequest{Ids: ids})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user details"})
			return
		}

		c.JSON(http.StatusOK, resp)
	})

	r.Run(":8080") // Start HTTP server on port 8080
}

// protoc --go_out=. --go-grpc_out=. user.proto
func main() {
	go startGRPCServer()
	startGinServer()
}
