package handler

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/shobky/giggs/pkg/_websocket"
)

func ServeWS(pool *_websocket.Pool, c *websocket.Conn) {
	fmt.Println(c.Locals("Host")) // "Localhost:3000"

	client := &_websocket.Client{
		Conn: c,
		Pool: pool,
		ID:   len(pool.Clients),
	}

	pool.Register <- client
	client.Read()
}
