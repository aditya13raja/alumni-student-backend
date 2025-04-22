package controllers

import (
	"context"
	"strings"
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTopic(c *fiber.Ctx) error {
	// Parse the request
	var req models.Topics
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	Topic := strings.ToLower(req.TopicName)

	// Check if topic already exists
	var existingTopic models.Topics

	// Mongodb query to check if topic exists
	err = utils.TopicsCollection.FindOne(
		context.Background(),
		bson.M{
			"topic_name": Topic,
		},
	).Decode(&existingTopic)

	// If topic exist give error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Topic already exists",
		})
	}

	// Create topic
	topic := models.Topics{
		ID:            primitive.NewObjectID(),
		TopicName:     Topic,
		Category:      req.Category,
		TopicFullName: req.TopicFullName,
		CreatedAt:     time.Now(),
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

func GetCategoryTopics(c *fiber.Ctx) error {
	category := c.Query("category")
	if strings.TrimSpace(category) == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing or invalid category parameter",
		})
	}

	var topics []*models.Topics
	cursor, err := utils.TopicsCollection.Find(context.Background(), bson.M{
		"category": category,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch topics",
		})
	}
	defer cursor.Close(context.Background())

	// Save all topics of the requested category
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
