package users

import (
	"context"
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

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}
func GetUserByIds() gin.HandlerFunc {
	return func(c *gin.Context) {
		idsStr := c.Request.FormValue("ids")
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
	}
}
