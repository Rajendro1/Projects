package httpgin

import (
	"github.com/gin-gonic/gin"
	"main.go/httpGin/users"
)

func StartServer() {
	r := gin.Default()

	r.GET("/user", users.GetUserById())
	r.GET("/users", users.GetUserByIds())

	r.Run(":8080")
}
