package controllers

import "github.com/gofiber/fiber/v2"

func CreateTopic(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "working",
	})
}
