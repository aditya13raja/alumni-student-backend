package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Topics struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TopicName     string             `bson:"topic_name" json:"topic_name"`
	TopicFullName string             `bson:"topic_fullname" json:"topic_fullname"`
	Category      string             `bson:"category" json:"category"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}
