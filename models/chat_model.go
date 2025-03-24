package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Topic          string             `bson:"topic" json:"topic"`
	Username       string             `bson:"username" json:"username"`
	Role           UserRole           `bson:"role" json:"role"`
	MessageContent string             `bson:"msg_content" json:"msg_content"`
	Timestamp      time.Time          `bson:"timestamp", json:"timestamp"`
}
