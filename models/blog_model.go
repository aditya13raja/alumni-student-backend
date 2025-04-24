package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blogs struct {
	ID         primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	Heading    string                 `bson:"heading" json:"heading"`
	Username   string                 `bson:"username" json:"username"`
	CoverImage string                 `bson:"cover_image" json:"cover_image"`
	Content    map[string]interface{} `bson:"content" json:"content"`
	CreatedAt  time.Time              `bson:"created_at,omitempty" json:"created_at"`
}
