package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Category            string             `bson:"category" json:"category"`
	CategoryFullName    string             `bson:"category_fullname" json:"category_fullname"`
	CategoryDescription string             `bson:"category_description" json:"category_description"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at"`
}
