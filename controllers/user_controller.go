package controllers

import (
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

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
	err := utils.UserCollection.FindOne(
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

// ------------------------------Update Profile--------------------------------
func UpdateUserProfile(c *fiber.Ctx) error {
	// Get username from the URL
	username := c.Params("username")

	// Parse the request body into a user struct
	var updatedData models.User
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Only allow updatable fields (for security)
	update := bson.M{
		"first_name":   updatedData.FirstName,
		"last_name":    updatedData.LastName,
		"email":        updatedData.Email,
		"age":          updatedData.Age,
		"degree":       updatedData.Degree,
		"major":        updatedData.Major,
		"passing_year": updatedData.PassingYear,
		"updated_at":   time.Now(),
	}

	// Update user in database
	result, err := utils.UserCollection.UpdateOne(
		context.Background(),
		bson.M{"username": username},
		bson.M{"$set": update},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user profile",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}

func DeleteUserProfile(c *fiber.Ctx) error {
	// Get username from URL
	username := c.Params("username")

	// Delete the user from database
	result, err := utils.UserCollection.DeleteOne(
		context.Background(),
		bson.M{"username": username},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user profile",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Profile deleted successfully",
	})
}
