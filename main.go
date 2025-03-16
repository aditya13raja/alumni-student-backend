package main

import (
	"log"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/routes"

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

	//---------------------------- PusherDB --------------------------
	configs.InitPusher()

	//---------------------------- Routes ------------------------------
	// Routes for auth
	routes.AuthRoutes(app)

	// Routes for chat

	// get port number
	Port := configs.GetPort()

	// Start server
	log.Fatal(app.Listen(":" + Port))

}
