package controllers

import (
	"goBackend/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "John",
	}

	user.LastName = "Doe"

	return c.JSON(user)
}
