package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shobky/giggs/api/router"
	"github.com/shobky/giggs/config/database"
)

func init() {
	database.Connect()
	database.Sync()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, *",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	router.Auth(app)
	router.User(app)
	app.Listen(":6969")

}
