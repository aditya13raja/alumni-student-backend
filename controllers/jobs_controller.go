package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/aditya13raja/alumni-student-backend/models"
	"github.com/aditya13raja/alumni-student-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateJobs(c *fiber.Ctx) error {
	var req *models.Jobs

	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error parsing jobs request",
		})
	}

	job := models.Jobs{
		ID:             primitive.NewObjectID(),
		Salary:         req.Salary,
		Username:       req.Username,
		JobRole:        req.JobRole,
		CompanyName:    req.CompanyName,
		Location:       req.Location,
		JobType:        req.JobType,
		JobMode:        req.JobMode,
		Validity:       req.Validity,
		JobLink:        req.JobLink,
		JobDescription: req.JobDescription,
		CreatedAt:      time.Now(),
	}

	isAlumni, err := utils.IsAlumni(job.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if !isAlumni {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authorized",
		})
	}

	fmt.Println(isAlumni)

	_, err = utils.JobsCollection.InsertOne(context.Background(), job)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error saving the job",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id": job.ID,
	})

}

func GetJobById(c *fiber.Ctx) error {
	idParams := c.Params("id")

	id, err := primitive.ObjectIDFromHex(idParams)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id parameter",
		})
	}

	var job models.Jobs
	err = utils.JobsCollection.FindOne(
		context.Background(),
		bson.M{
			"_id": id,
		},
	).Decode(&job)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Job not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"job": job,
	})
}

func GetJobsList(c *fiber.Ctx) error {
	ctx := context.TODO()
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetProjection(bson.M{
		"job_role":     1,
		"company_name": 1,
		"location":     1,
		"job_type":     1,
		"job_mode":     1,
		"job_link":     1,
		"created_at":   1,
		"validity":     1,
	})

	cursor, err := utils.JobsCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	var jobs []models.Jobs
	if err := cursor.All(ctx, &jobs); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to parse jobs"})
	}

	return c.JSON(jobs)
}
