package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/api/handler"
	"github.com/shobky/giggs/api/middleware"
)

func User(app *fiber.App) {
	r := app.Group("/user")

	r.Get("/user", middleware.DeserializeUser, handler.GetCurrentUser)
}
