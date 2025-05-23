package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	//	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func AlumniRoutes(app *fiber.App) {
	alumni := app.Group("/api/alumni")
	//alumni := app.Group("/api/alumni", utils.AuthMiddleware)

	alumni.Get("/list", controllers.GetAllAlumni)
}
