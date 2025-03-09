package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "test route is working correctly!"})
	})

	auth.Post("/signup", controllers.SignUp)
	auth.Post("/signin", controllers.SignIn)
	auth.Get("/signout", controllers.SignOut)
}
