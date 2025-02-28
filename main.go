package main

import (
	"alumni-student-backend/configs"
	"alumni-student-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Create new instance of fiber
	app := fiber.New()

	//---------------------------- GetEnv ---------------------------
	// Get .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	//---------------------------- Database ---------------------------
	// Connect mongodb database
	configs.ConnectDB()
	defer configs.DisconnectDB()

	//---------------------------- Routes ------------------------------
	// Routes for auth
	routes.AuthRoutes(app)

	// get port number
	Port := configs.GetPort()

	// Start server
	log.Fatal(app.Listen(":" + Port))

}
