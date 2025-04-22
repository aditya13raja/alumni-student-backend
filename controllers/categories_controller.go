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

func CreateCategory(c *fiber.Ctx) error {
	// Parse the request
	var req models.Categories
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	Category := strings.ToLower(req.Category)

	// Check if category already exists
	var existingCategory models.Categories

	// Mongodb query to check if topic exists
	err = utils.CategoriesCollection.FindOne(
		context.Background(),
		bson.M{
			"category": Category,
		},
	).Decode(&existingCategory)

	// If category exist give error
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Category already exists",
		})
	}

	// Create topic
	category := models.Categories{
		ID:                  primitive.NewObjectID(),
		Category:            Category,
		CategoryFullName:    req.CategoryFullName,
		CategoryDescription: req.CategoryDescription,
		CreatedAt:           time.Now(),
	}

	//Create topic in database
	_, err = utils.CategoriesCollection.InsertOne(context.Background(), category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error creating category"})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"category": category,
	})
}

func GetAllCategories(c *fiber.Ctx) error {
	var categories []*models.Categories

	// Save all categories to the categories array(slice) created
	cursor, err := utils.CategoriesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	for cursor.Next(context.Background()) {
		var category *models.Categories
		if err := cursor.Decode(&category); err == nil {
			categories = append(categories, category)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}
