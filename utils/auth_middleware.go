package utils

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from cookie
	token := c.Cookies("access_token")

	// Check if token is empty
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized, missing token",
		})
	}

	// Validate token
	userID, err := ValidateJWT(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized, invalid token",
		})
	}

	// Store userID in fiber's Locals (for access in other handlers)
	c.Locals("userID", userID)

	return c.Next()
}
