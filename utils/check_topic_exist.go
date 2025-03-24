package utils

import (
	"context"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Check if topic already exists
func CheckTopicExists(c *fiber.Ctx, topic string) error {
	// Create a variable to store topic
	var existingTopic *models.Topics

	// Mongodb query to check if topic exists
	err := configs.TopicsCollection.FindOne(
		context.Background(),
		bson.M{
			"topic_name": topic,
		},
	).Decode(&existingTopic)

	// If topic exist give error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Topic already exists",
		})
	}

	// If topic doesn't exist don't return any error
	return nil
}
