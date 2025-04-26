package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func ChatRoutes(app *fiber.App) {
	chat := app.Group("/api/chat", utils.AuthMiddleware)

	chat.Post("/send", controllers.SendMessage)
	chat.Get("/get/:topic", controllers.GetMessageByTopic)
}
