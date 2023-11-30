package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn     *websocket.Conn
	username string
}

var clients = make(map[*websocket.Conn]Client)
var broadcast = make(chan Message)

func handleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	username := c.Query("username")
	client := Client{conn: ws, username: username}
	clients[ws] = client

	fmt.Printf("New connection: %s\n", username)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			delete(clients, ws)
			break
		}

		// Broadcast the received message to all clients
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		// Send the message to all connected clients
		for _, client := range clients {
			err := client.conn.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Error writing message: %v\n", err)
				delete(clients, client.conn)
				break
			}
		}
	}
}

// NewMessageHandler handles incoming messages from the clients
func NewMessageHandler(c *gin.Context) {
	var msg Message
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}

	// Broadcast the received message to all clients
	broadcast <- msg

	c.JSON(http.StatusOK, gin.H{"status": "Message sent successfully"})
}

func main() {
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
