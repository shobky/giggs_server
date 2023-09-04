package router

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/pkg/_websocket"
)

func Websocket(app *fiber.App) {

	pool := _websocket.NewPool()
	go pool.Start()

	r := app.Group("/ws")

	r.Use("/", func(c *fiber.Ctx) error {
		return c.Next()
	})

	r.Get("/", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"

		client := &_websocket.Client{
			Conn: c,
			Pool: pool,
			ID:   len(pool.Clients),
		}

		pool.Register <- client
		client.Read()
	}))
}
