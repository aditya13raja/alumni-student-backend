package controllers

import (
	"context"
	"fmt"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllAlumni(c *fiber.Ctx) error {
	var alumni []*models.User

	// Project only the required fields
	projection := bson.M{
		"first_name":   1,
		"last_name":    1,
		"email":        1,
		"username":     1,
		"age":          1,
		"passing_year": 1,
		"degree":       1,
		"major":        1,
		"_id":          0,
	}

	findOptions := options.Find().SetProjection(projection)

	// Save all Alumni to the categories array(slice) created
	cursor, err := utils.UserCollection.Find(context.Background(), bson.M{"role": "alumni"}, findOptions)
	if err != nil {
		fmt.Print(err)
		return err
	}

	if err := cursor.All(context.Background(), &alumni); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse alumni data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"alumni": alumni,
	})
}
