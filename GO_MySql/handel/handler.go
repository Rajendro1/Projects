package handel

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/api"
	"main.go/includes"
)

// CORS Middleware
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

func HandleRequest() {
	log.Println("Before Open  This Server  Run This Command in server 'sudo ufw allow 1010'")
	log.Println("Quit The Server With CONTROL-C.")
	r := gin.Default()
	r.Use(CORS)

	r.POST("/User_Authentication", api.PostUsers)
	r.GET("/All_users_By_Super_Admin", api.GetUsers)
	r.POST("/SuperAdmin_Authentication", api.PostSuperAdmin)

	r.POST("/Product", api.PostProduct)
	r.PUT("/Product", api.UpdateProduct)
	r.GET("/Product", api.GetProduct)
	r.DELETE("/Product", api.DeleteProduct)

	r.GET("/Users_Activity", api.GetUsersActivity)
	//Local
	// r.Run(includes.DB_HOST + ":" + includes.SERVER_PORT)

	// Server
	// r.Run(includes.SERVER_HOST + ":" + includes.SERVER_PORT)
	r.RunTLS(":"+includes.SERVER_PORT, includes.SERVER_CERT_FILE, includes.SERVER_KEY_FILE)
}
