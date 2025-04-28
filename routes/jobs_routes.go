package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func JobsRoutes(app *fiber.App) {
	jobs := app.Group("/api/jobs", utils.AuthMiddleware)

	jobs.Post("/create", controllers.CreateJobs)
	jobs.Get("/:id", controllers.GetJobById)
	jobs.Get("/list/jobs", controllers.GetAllJobs)
	jobs.Get("/latest/jobs", controllers.GetLatestJobs)
}
