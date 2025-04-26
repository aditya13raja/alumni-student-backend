package utils

import (
	"context"
	"strings"

	"github.com/aditya13raja/alumni-student-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func IsAlumni(username string) (bool, error) {
	var user models.User

	err := UserCollection.FindOne(
		context.Background(),
		bson.M{"username": username},
	).Decode(&user)

	if err != nil {
		return false, err
	}

	return strings.ToLower(strings.TrimSpace(string(user.Role))) == "alumni", nil
}
