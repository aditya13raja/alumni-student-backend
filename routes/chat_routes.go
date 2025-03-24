package routes

import (
	"github.com/aditya13raja/alumni-student-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ChatRoutes(app *fiber.App) {
	chat := app.Group("/api/chat")

	chat.Post("/send", controllers.SendMessage)
}
