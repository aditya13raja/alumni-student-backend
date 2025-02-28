package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole string

const (
	Alumni  UserRole = "alumni"
	Student UserRole = "student"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName   string             `json:"first_name" bson:"first_name"`
	LastName    string             `json:"last_name" bson:"last_name"`
	Email       string             `json:"email" bson:"email"`
	Age         int                `json:"age" bson:"age"`
	Role        UserRole           `json:"role" bson:"role"`
	Degree      string             `json:"degree" bson:"degree"`
	Major       string             `json:"major" bson:"major"`
	PassingYear int                `json:"passing_year,omitempty" bson:"passing_year,omitempty"`
	Username    string             `json:"user_name" bson:"user_name"`
	Password    string             `json:"password" bson:"password"`
	Updated_at  time.Time          `json:"updated_at" bson:"updated_at"`
	Created_at  time.Time          `json:"created_at" bson:"created_at"`
}
