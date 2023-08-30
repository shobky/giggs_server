package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/pkg/user"
)

func GetCurrentUser(c *fiber.Ctx) error {
	emailValue := c.Locals("email")
	if emailValue == "" {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't get this profile",
			"message": "Something went wrong, try to login again",
		})
	}

	email, ok := emailValue.(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't get this profile",
			"message": "Something went wrong, try to login again",
		})
	}

	user, ok := user.FindUserByEmail(email)
	if !ok {
		return c.Status(404).JSON(fiber.Map{
			"title":   "Unauthroized",
			"message": "sorry, we cound't find you account.",
			"error":   fmt.Sprintf("user %s not found", email),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": fmt.Sprintf("Welocme %s.", user.GivenName),
		"user":    user,
	})
}
