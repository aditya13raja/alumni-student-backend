package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetAuthCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // one week expiry
		HTTPOnly: true,
		Secure:   false, //TODO: set to true at time of hosting
		SameSite: "Strict",
	})
}

func ClearAuthCookie(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 1),
		HTTPOnly: true,
		Secure:   false, //TODO: set to true at time of hosting
		SameSite: "Strict",
	})
}
