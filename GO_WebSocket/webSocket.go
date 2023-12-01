package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrader is used to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

// handleWebSocket handles new WebSocket connections using Gin.
func handleWebSocket1(c *gin.Context) {
	// Upgrade the HTTP connection to a WebSocket connection
	// conn1, err1 := upgrader.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New WebSocket connection")

	// Handle messages from this connection
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", message)

		// Echo the message back to the client
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}
}

func websocketv1() {
	// Create a Gin router
	r := gin.Default()

	// Define a route for WebSocket connections
	r.GET("/ws", handleWebSocket1)

	// Start the Gin server on port 8080
	fmt.Println("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Run error:", err)
	}
}
