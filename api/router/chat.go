package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/api/handler"
	"github.com/shobky/giggs/api/middleware"
)

func Chat(app *fiber.App) {
	r := app.Group("/chat")

	r.Use(middleware.DeserializeUser)

	// :id is the id of the user to create a chat with
	r.Get("/", handler.GetAllContacts)
	r.Get("/new/:userID", handler.NewContact)
	r.Get("/:id", handler.GetContact)
}
