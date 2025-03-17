package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	protected := app.Group("/api/protected", middleware.AuthMiddleware)

	protected.Get("/:username", controllers.GetUserProfile)
}
