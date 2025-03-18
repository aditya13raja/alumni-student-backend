package controllers

import (
	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"

	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// -------------------------- Sign Up controller ---------------------------------
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
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Database Error",
			})
		}
	} else { // Found existing user
		if existingUser.Email == req.Email {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Email already in use",
			})
		}
		if existingUser.Username == req.Username {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Username already exists",
			})
		}
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create user
	user := models.User{
		ID:          primitive.NewObjectID(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Age:         req.Age,
		Role:        req.Role,
		Degree:      req.Degree,
		Major:       req.Major,
		Username:    req.Username,
		PassingYear: req.PassingYear,
		Password:    string(hashedPassword),
	}

	// Add user to mongodb
	_, err = configs.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate JWT token",
		})
	}

	// Send token as cookies
	utils.SetAuthCookie(c, token)

	// Return token
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    utils.FormatUserResponse(user),
	})
}

// ----------------------------- Sign In Controller -------------------------------
func SignIn(c *fiber.Ctx) error {
	// Parse the user request body(json)
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find the user with requested username
	var user models.User
	err := configs.UserCollection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Username not found",
		})
	}

	// Check if password is correct or not
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Password is wrong",
		})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate JWT token",
		})
	}

	// Send token as cookies
	utils.SetAuthCookie(c, token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successfull",
		"user":    utils.FormatUserResponse(user),
	})
}

// ------------------------------Sign Out------------------------------------
func SignOut(c *fiber.Ctx) error {
	utils.ClearAuthCookie(c)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successfull",
	})
}

// ----------------------------- Test ---------------------------
func Test(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Test working",
	})
}
