package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/api/schema"
)

func CreateSession(c *fiber.Ctx, ID uint, Email string) error {
	tokenString, tokenError := GenerateToken(&schema.TokenPayload{
		ID:    ID,
		Email: Email,
	})
	if tokenError != nil {
		return tokenError
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accesstoken",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 14),
		SameSite: "None",
		Secure:   false,
		Path:     "/",
		HTTPOnly: true,
	})

	return nil
}
