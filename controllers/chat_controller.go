package controllers

import (
	"context"
	"time"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
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
	_, err := utils.ChatsCollection.InsertOne(context.Background(), chat)

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

// -------------------------------- GetMessageByTopic --------------------------------
func GetMessageByTopic(c *fiber.Ctx) error {
	// Get topic name from params
	topic := c.Params("topic")

	// Create slice of chats to store all chats
	var chats []models.Chat

	// Get all chats in ascending order from mondodb
	cursor, err := utils.ChatsCollection.Find(
		context.Background(),
		bson.M{"topic": topic},
		// &options.FindOptions{
		// 	Sort: bson.D{{"timestamp", 1}},
		// },
	)

	if err != nil {
		utils.HandleError(c, fiber.StatusInternalServerError, "Failed to get messages")
	}

	// Iterate and save chats to messages variable
	for cursor.Next(context.Background()) {
		var chat models.Chat
		if err := cursor.Decode(&chat); err == nil {
			chats = append(chats, chat)
		}
	}

	// If Chats is empty
	if len(chats) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No message found for this topic",
		})
	}

	// Return all chat messages
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"chats": chats,
	})
}
