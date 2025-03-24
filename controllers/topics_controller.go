package controllers

import (
	"context"
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTopic(c *fiber.Ctx) error {
	// Parse the request
	var req *models.Topics
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if topic already exists
	if err := utils.CheckTopicExists(c, req.TopicName); err != nil {
		return err
	}

	// Create topic
	topic := models.Topics{
		ID:        primitive.NewObjectID(),
		TopicName: req.TopicName,
		CreatedAt: time.Now(),
	}

	//Create topic in database
	_, err = utils.TopicsCollection.InsertOne(context.Background(), topic)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error creating topic"})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"topic": topic,
	})
}

func GetAllTopics(c *fiber.Ctx) error {
	var topics []*models.Topics

	// Save all topics to the topics array(slice) created
	cursor, err := utils.TopicsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	for cursor.Next(context.Background()) {
		var topic *models.Topics
		if err := cursor.Decode(&topic); err == nil {
			topics = append(topics, topic)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"topics": topics,
	})
}
