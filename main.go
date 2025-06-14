package main

import (
	"log"
	"os"

	"github.com/aditya13raja/alumni-student-backend/configs"
	"github.com/aditya13raja/alumni-student-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Create new instance of fiber
	app := fiber.New()

	//---------------------------- GetEnv ---------------------------
	// Load .env file only in local development
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file:", err)
		}
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

	// Routes for user profile
	routes.UserRoutes(app)

	// Routes for chat
	routes.ChatRoutes(app)

	// Routes for topics
	routes.TopicsRoutes(app)

	// Routes for category
	routes.CategoryRoutes(app)

	// Routes for blog
	routes.BlogRoutes(app)

	// Routes for jobs
	routes.JobsRoutes(app)

	// Routes for alumni
	routes.AlumniRoutes(app)

	// get port number
	Port := configs.GetPort()

	// Start server
	log.Fatal(app.Listen(":" + Port))

}
