package controllers

import (
	"context"
	"time"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ------------------------- Send Message --------------------------
func SendMessage(c *fiber.Ctx) error {
	var chat models.Chat

	// Parse the req body into chat struct
	if err := c.BodyParser(&chat); err != nil {
		utils.HandleError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Create id and timestamp for chat created
	chat.ID = primitive.NewObjectID()
	chat.Timestamp = time.Now()

	// Save chat to mongodb chat collection
	_, err := configs.ChatsCollection.InsertOne(context.Background(), chat)

	if err != nil {
		utils.HandleError(c, fiber.StatusBadRequest, "Failed to send message")
	}

	// Trigger pusher event to broadcast messages in real-time
	err = configs.PusherClient.Trigger(chat.Topic, "new-event", chat)

	if err != nil {
		utils.HandleError(c, fiber.StatusInternalServerError, "Failed to trigger pusher event")
	}

	return c.Status(fiber.StatusOK).JSON(chat)
}
