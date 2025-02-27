package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() {
	// get uri from .env file
	MongoDB_URI := os.Getenv("MONGODB_URI")
	if MongoDB_URI == "" {
		log.Fatal("Error loading MongoDB_URI from .env")
	}

	clientOptions := options.Client().ApplyURI(MongoDB_URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	CheckError(err)

	err = client.Ping(context.Background(), nil)
	CheckError(err)

	fmt.Println("Connected to MongoDB ✅")
}
