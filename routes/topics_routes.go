package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func TopicsRoutes(app *fiber.App) {
	topics := app.Group("/api/topics")

	topics.Post("/create-topic", controllers.CreateTopic)
}
