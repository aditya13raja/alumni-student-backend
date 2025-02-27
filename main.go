package main

import (
	"alumni-student-backend/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Create new instance of fiber
	app := fiber.New()

	// Get .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Connect mongodb database
	configs.ConnectDB()

	// Check server running
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(201).SendString("hello")
	})

	// get port number
	Port := configs.GetPort()

	// Start server
	log.Fatal(app.Listen(":" + Port))

}
