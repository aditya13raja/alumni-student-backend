package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func BlogRoutes(app *fiber.App) {
	blog := app.Group("/api/blog", middleware.AuthMiddleware)

	blog.Post("/save-blog", controllers.SaveBlog)
	blog.Get("/:id", controllers.GetBlogById)
	blog.Get("/list/blogs", controllers.GetBlogsList)
}
