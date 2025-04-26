package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	profile := app.Group("/api/user", utils.AuthMiddleware)

	profile.Get("/:username", controllers.GetUserProfile)
}
