package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shobky/giggs/api/schema"
	"github.com/shobky/giggs/model"
	"github.com/shobky/giggs/pkg/auth"
	"github.com/shobky/giggs/pkg/user"
	"github.com/shobky/giggs/utils"
)

func Signup(c *fiber.Ctx) error {
	var body schema.SignupBody
	if err := utils.ParseBodyAndValidate(c, &body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Invalid form field",
			"message": err.Message,
		})
	}

	// ok is true if a user is found
	_, ok := user.FindUserByEmail(body.Email)
	if ok {
		return c.Status(409).JSON(fiber.Map{
			"title":   "Account already exists.",
			"message": "As long as we would love it we can't have two of you.",
		})
	}

	hash, hashErr := auth.HashPassword(body.Password)
	if hashErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"title":   "Failed to signup.",
			"message": "Something wrong with your password, please try again.",
		})
	}

	data := &model.User{
		Email:         body.Email,
		EmailVerified: false,
		Password:      hash,
		GivenName:     body.GivenName,
		FamilyName:    body.FamilyName,
		Name:          body.Name,
		Provider:      "local",
	}

	err := user.InsertNew(data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't create your account.",
			"message": "something went wrong, try using another email adress",
			"error":   err.Error(),
		})
	}

	seshErr := auth.CreateSession(c, data.ID, data.Email)
	if seshErr != nil {
		c.Status(500).JSON(fiber.Map{
			"title":   "Failed to signup",
			"message": "something went wrong, can't sign user in",
			"error":   seshErr.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"user": data,
	})
}

func Signin(c *fiber.Ctx) error {
	var body schema.LoginBody
	if err := utils.ParseBodyAndValidate(c, &body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Invalid form field",
			"message": err.Message,
		})
	}

	// ok is true if a user is found
	user, ok := user.FindUserByEmail(body.Email)
	if !ok {
		return c.Status(404).JSON(fiber.Map{
			"title":   "We can't find you.",
			"message": "This email address is not registered, please sign up.",
		})
	}

	err := auth.ComparePassword(user.Password, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"title":   "wrong password",
			"message": fmt.Sprintf("wrong password for %s", user.Email),
			"error":   err.Error(),
		})
	}

	seshErr := auth.CreateSession(c, user.ID, user.Email)
	if seshErr != nil {
		c.Status(500).JSON(fiber.Map{
			"title":   "Failed to signup",
			"message": "something went wrong, can't sign user in",
			"error":   seshErr.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

func Signout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "accesstoken",
		Expires: time.Now().Add(-time.Hour * 24 * 14),
	})
	return c.Status(200).JSON(fiber.Map{
		"message": "signed out",
	})
}

func UseGoogle(c *fiber.Ctx) error {

	var request schema.CallbackToken
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't continue.",
			"message": "something wen't wrong please try again.",
			"error":   err.Error(),
		})
	}

	payload, parseErr := auth.ParseIdtoken(request.Token)
	if parseErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"title":   "Can't continue.",
			"message": "something wen't wrong please try again.",
			"error":   parseErr.Error(),
		})
	}

	var data = &model.User{
		Email:         payload.Claims["email"].(string),
		EmailVerified: payload.Claims["email_verified"].(bool),
		Name:          payload.Claims["name"].(string),
		GivenName:     payload.Claims["given_name"].(string),
		FamilyName:    payload.Claims["family_name"].(string),
		Picture:       payload.Claims["picture"].(string),
		Provider:      "google",
	}

	_, ok := user.FindUserByEmail(data.Email)
	if !ok {
		err := user.InsertNew(data)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"title":   "something wen't wrong",
				"message": "please try using another email adress",
				"error":   err.Error(),
			})
		}
	}

	seshErr := auth.CreateSession(c, data.ID, data.Email)
	if seshErr != nil {
		c.Status(500).JSON(fiber.Map{
			"title":   "Failed to signup",
			"message": "something went wrong, can't sign user in",
			"error":   seshErr.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"user": data,
	})
}
