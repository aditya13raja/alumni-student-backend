package controllers

import (
	"context"
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveBlog(c *fiber.Ctx) error {
	var req models.Blogs
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse blog request",
		})
	}

	blog := models.Blogs{
		ID:         primitive.NewObjectID(),
		Heading:    req.Heading,
		Username:   req.Username,
		CoverImage: req.CoverImage,
		Content:    req.Content,
		CreatedAt:  time.Now(),
	}

	_, err = utils.BlogsCollection.InsertOne(context.Background(), blog)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Creating Blog",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": blog.ID,
	})
}
