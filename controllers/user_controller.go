package controllers

import (
	"alumni-student-backend/configs"
	"alumni-student-backend/models"
	"alumni-student-backend/utils"

	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	// Parse request body
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Check if email or username already exists
	var existingUser models.User
	err := configs.UserCollection.FindOne(
		context.Background(),
		bson.M{"$or": []bson.M{
			{"email": req.Email},
			{"username": req.Username},
		}},
	).Decode(&existingUser)

	// If found existing email or username
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No existing user found, continue with registration
		} else {
			fmt.Println("Error from FindOne:", err) // Debugging output
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database Error"})
		}
	} else { // Found existing user
		fmt.Println("Found existing user:", existingUser) // Debugging output
		if existingUser.Email == req.Email {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already in use"})
		}
		if existingUser.Username == req.Username {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
		}
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// Create user
	user := models.User{
		ID:        primitive.NewObjectID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Age:       req.Age,
		Role:      req.Role,
		Degree:    req.Degree,
		Major:     req.Major,
		Username:  req.Username,
		Password:  string(hashedPassword),
	}

	// Add user to mongodb
	_, err = configs.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT token"})
	}

	// Return token
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully", "token": token})
}
