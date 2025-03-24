package utils

import "github.com/gofiber/fiber/v2"

func HandleError(c *fiber.Ctx, statusCode int, errorMessage string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": errorMessage,
	})
}
