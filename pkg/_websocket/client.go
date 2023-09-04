package _websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	ID   int
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type     int    `json:"type"`
	Body     string `json:"body"`
	SenderID int    `json:"sender_id"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()
	connection := c.Conn

	for {
		mt, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		message := Message{Type: mt, Body: string(msg), SenderID: c.ID}
		c.Pool.Brodcast <- message
		fmt.Printf("msg rcvd:%+v\n", message)
	}
}
