package configs

import (
	"context"
	"log"
	"os"

	"github.com/aditya13raja/alumni-student-backend/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create client for connecting to DB
var (
	client *mongo.Client
	DB     *mongo.Database
)

// Create Collection for user
var (
	UserCollection   *mongo.Collection
	TopicsCollection *mongo.Collection
	ChatsCollection  *mongo.Collection
)

func ConnectDB() {
	// get uri from .env file
	MongoDB_URI := os.Getenv("MONGODB_URI")
	if MongoDB_URI == "" {
		log.Fatal("Error loading MongoDB_URI from .env")
	}

	clientOptions := options.Client().ApplyURI(MongoDB_URI)

	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	utils.CheckError(err)

	err = client.Ping(context.Background(), nil)
	utils.CheckError(err)

	log.Println("Connected to MongoDB ✅")

	// Create database named alumni-student-db in mongodb database
	DB = client.Database("alumni-student-db")

	// Create collections for storing different types of data
	UserCollection = DB.Collection("users")
	TopicsCollection = DB.Collection("topics")
	ChatsCollection = DB.Collection("chats")
}

func DisconnectDB() {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Println("Error disconnecting mongodb: ", err)
		} else {
			log.Println("Disconnected to MongoDB ✅")
		}
	}
}
