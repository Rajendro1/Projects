package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(r, &http2.Server{}),
	}

	log.Printf("Listening on http://localhost:8080\n")
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
