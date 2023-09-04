package middleware

import (
	"fmt"

	"github.com/shobky/giggs/pkg/auth"

	"github.com/gofiber/fiber/v2"
)

func DeserializeUser(c *fiber.Ctx) error {

	tokenString := c.Cookies("accesstoken")
	fmt.Printf("token: %s", tokenString)
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"title":   "Unauthroized",
			"message": "You need to login first",
			"error":   "token is empty",
		})
	}

	payload, err := auth.Verify(tokenString)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"title":   "Unauthroized",
			"message": "You need to login first",
			"error":   err.Error(),
		})
	}

	c.Locals("email", payload.Email)
	c.Locals("id", payload.ID)

	return c.Next()

}
