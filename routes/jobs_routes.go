package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func JobsRoutes(app *fiber.App) {
	jobs := app.Group("/api/jobs")

	jobs.Post("/create", controllers.CreateJobs)
	jobs.Get("/:id", controllers.GetJobById)
	jobs.Get("/list/jobs", controllers.GetJobsList)
}
