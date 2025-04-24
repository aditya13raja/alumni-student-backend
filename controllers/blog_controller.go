package controllers

import (
	"context"
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetBlogById(c *fiber.Ctx) error {
	idParams := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id parameter",
		})
	}

	var blog models.Blogs
	err = utils.BlogsCollection.FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	).Decode(&blog)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Blog not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blog": blog,
	})
}

func GetBlogsList(c *fiber.Ctx) error {
	ctx := context.TODO()
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetProjection(bson.M{
		"heading":     1,
		"username":    1,
		"cover_image": 1,
		"created_at":  1,
	})

	cursor, err := utils.BlogsCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	var blogs []models.Blogs
	if err := cursor.All(ctx, &blogs); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse blogs"})
	}

	return c.JSON(blogs)
}
