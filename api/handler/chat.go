package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/model"
	"github.com/shobky/giggs/pkg/chat"
	"github.com/shobky/giggs/utils"
)

func NewContact(c *fiber.Ctx) error {

	// don't forget this needs more error handling and security addition........ ------+++++++
	id1, ok := c.Locals("id").(uint)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request: Invalid user IDs.",
		})
	}
	user2val := c.Params("userID")
	if user2val == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request: Invalid user IDs.",
		})
	}
	id2 := utils.StringToUint(user2val)

	data := &model.Contact{
		UserID:        id1,
		ContactUserID: id2,
	}

	err := chat.InsertNewContact(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't create this contact room.",
			"message": "something went wrong, try again later.",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"contact": data,
	})
}

func GetContact(c *fiber.Ctx) error {
	contactIDval := c.Params("id")
	ContactID := utils.StringToUint(contactIDval)

	data, ok := chat.GetContacgByID(ContactID)
	if !ok {
		return c.Status(404).JSON(fiber.Map{
			"message": "this contact doesn't exist",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"contact": data,
	})

}

func GetAllContacts(c *fiber.Ctx) error {
	// gets user id from the middleware
	userID, ok := c.Locals("id").(uint)
	if !ok {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	contacts, err := chat.GetAllContacts(userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "no contacts",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"contacts": contacts,
	})

}
