package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// ------------------------------User Profile--------------------------------
func UserProfile(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"user": "successfull",
	})
}
