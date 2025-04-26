package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func TopicsRoutes(app *fiber.App) {
	topics := app.Group("/api/topics", utils.AuthMiddleware)

	topics.Post("/create-topic", controllers.CreateTopic)
	topics.Get("/get-topics", controllers.GetAllTopics)
	topics.Get("/", controllers.GetCategoryTopics)

}
