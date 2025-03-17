package controllers

import (
	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"

	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ------------------------------User Profile--------------------------------
func GetUserProfile(c *fiber.Ctx) error {
	// Get username from the url
	username := c.Params("username")

	// Return the user details
	// Create a user variable to store user details
	var user models.User

	// Find the username in database
	err := configs.UserCollection.FindOne(
		context.Background(),
		bson.M{"username": username},
		// Hide password
		options.FindOne().SetProjection(bson.M{"password": 0}),
	).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}
