package httpgin

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/test/bufconn"
	"main.go/httpGin/users"
)

func StartServer() {
	r := gin.Default()

	r.GET("/user", users.GetUserById(&bufconn.Listener{}))
	r.GET("/users", users.GetUserByIds())

	r.Run(":8080")
}
