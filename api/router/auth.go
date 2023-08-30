package router

import (
	"github.com/shobky/giggs/api/handler"
	"github.com/shobky/giggs/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/signup", handler.Signup)
	r.Post("/signin", handler.Signin)
	r.Post("/google/callback", handler.UseGoogle)

	r.Get("/signout", middleware.DeserializeUser, handler.Signout)
}
