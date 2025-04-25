package utils

import "go.mongodb.org/mongo-driver/mongo"

// Create Collection for user
var (
	UserCollection       *mongo.Collection
	TopicsCollection     *mongo.Collection
	ChatsCollection      *mongo.Collection
	CategoriesCollection *mongo.Collection
	BlogsCollection      *mongo.Collection
	JobsCollection       *mongo.Collection
)

// Create collections for storing different types of data
func CreateCollection(DB *mongo.Database) {
	UserCollection = DB.Collection("users")
	TopicsCollection = DB.Collection("topics")
	ChatsCollection = DB.Collection("chats")
	CategoriesCollection = DB.Collection("categories")
	BlogsCollection = DB.Collection("blogs")
	JobsCollection = DB.Collection("jobs")
}
