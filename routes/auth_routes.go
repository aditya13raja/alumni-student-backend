package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/signup", controllers.SignUp)
	auth.Post("/signin", controllers.SignIn)
	auth.Get("/signout", controllers.SignOut)
}
