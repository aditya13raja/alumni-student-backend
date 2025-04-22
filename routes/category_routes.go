package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	categories := app.Group("/api/categories")

	categories.Post("/create-category", controllers.CreateCategory)
	categories.Get("/get-categories", controllers.GetAllCategories)
}
