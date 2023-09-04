package utils

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) error {
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	fmt.Println("pares", body)

	return Validate(body)
}

func ToString(value interface{}) (string, *fiber.Error) {
	if value == "" {
		return "", fiber.NewError(400, "Value is empty.")
	}

	newString, ok := value.(string)
	if !ok {
		return "", fiber.NewError(400, "Something went wrong.")
	}

	return newString, nil
}

func StringToUint(value string) uint {
	ui64, _ := strconv.ParseUint(value, 10, 64)
	ui := uint(ui64)
	return ui
}

// GetUser is helper function for getting authenticated user's id
// func GetUser(c *fiber.Ctx) *uint {
// 	id, _ := c.Locals("USER").(uint)
// 	return &id
// }
