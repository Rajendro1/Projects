package httpgin_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"main.go/httpGin/users"
	userpb "main.go/userProto"
)

type mockUserServiceServer struct {
	userpb.UnimplementedUserServiceServer
}

func (s *mockUserServiceServer) GetUserById(ctx context.Context, req *userpb.UserRequest) (*userpb.User, error) {
	if req.Id == 1 {
		return &userpb.User{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true}, nil
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}
func TestGetUserById(t *testing.T) {
	// Setup the mock gRPC server
	lis := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &mockUserServiceServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()
	defer s.Stop()

	// Setup Gin HTTP server
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	r.GET("/user", users.GetUserById(lis))
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Define test cases
	tests := []struct {
		userId         string
		expectedStatus int
		expectedBody   string
	}{
		{"1", http.StatusOK, `{"id":1,"fname":"Steve","city":"LA","phone":1234567890,"height":5.8,"married":true}`},
		// {"2", http.StatusNotFound, `{"error":"User not found"}`},
		// {"abc", http.StatusBadRequest, `{"error":"Invalid user ID"}`},
	}

	// Run test cases
	for _, tc := range tests {
		res, err := http.Get(ts.URL + "/user?id=" + tc.userId)
		if err != nil {
			t.Fatalf("Could not send GET request: %v", err)
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Could not read response body: %v", err)
		}

		if res.StatusCode != tc.expectedStatus || string(body) != tc.expectedBody {
			t.Errorf("For user ID %v, expected (%v, %v), got (%v, %v)", tc.userId, tc.expectedStatus, tc.expectedBody, res.StatusCode, string(body))
		}
	}
}
