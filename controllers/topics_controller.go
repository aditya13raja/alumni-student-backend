package controllers

import (
	"context"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTopic(c *fiber.Ctx) error {
	// Parse the request
	var req *models.Topics
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Create topic
	topic := models.Topics{
		ID:        primitive.NewObjectID(),
		TopicName: req.TopicName,
	}

	//Create topic in database
	_, err = configs.TopicsCollection.InsertOne(context.Background(), topic)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error creating topic"})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"topic": topic,
	})
}
