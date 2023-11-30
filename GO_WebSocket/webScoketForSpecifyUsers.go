package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ClientV2 struct {
	conn     *websocket.Conn
	username string
}

var clientsV2 = make(map[string]Client)
var broadcastV2 = make(chan Message)

func handleConnectionsV2(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	username := c.Query("username")
	client := Client{conn: ws, username: username}
	clientsV2[username] = client

	fmt.Printf("New connection: %s\n", username)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			delete(clientsV2, username)
			break
		}

		// Broadcast the received message to all clients
		broadcast <- msg
	}
}

func handleMessagesV2() {
	for {
		msg := <-broadcast

		// Send the message to a specific user
		if client, ok := clientsV2[msg.Username]; ok {
			err := client.conn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Error writing message to %s: %v\n", msg.Username, err)
				delete(clientsV2, msg.Username)
			}
		}
	}
}

// NewMessageHandler handles incoming messages from the clients
func NewMessageHandlerV2(c *gin.Context) {
	var msg Message
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}

	// Broadcast the received message to a specific user
	broadcast <- msg

	c.JSON(http.StatusOK, gin.H{"status": "Message sent successfully"})
}

func webScokitV2() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		handleConnections(c)
	})

	r.POST("/send-message", NewMessageHandler)

	go handleMessages()

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
