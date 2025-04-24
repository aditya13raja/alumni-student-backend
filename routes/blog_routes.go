package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(app *fiber.App) {
	blog := app.Group("/api/blog")

	blog.Post("/save-blog", controllers.SaveBlog)
}
