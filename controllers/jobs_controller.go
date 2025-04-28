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

func fetchJobs(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]models.Jobs, error) {
	cursor, err := utils.JobsCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var jobs []models.Jobs
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, err
	}

	return jobs, nil
}

func GetAllJobs(c *fiber.Ctx) error {
	ctx := context.Background()

	// No limit or sorting for all jobs, just fetching them
	opts := options.Find()

	// Call fetchJobs to get all jobs
	jobs, err := fetchJobs(ctx, bson.M{}, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// Return the list of jobs as JSON
	return c.JSON(jobs)
}

func GetLatestJobs(c *fiber.Ctx) error {
	ctx := context.Background()

	// Sort by created_at in descending order and limit to 6
	opts := options.Find().SetSort(bson.D{{"created_at", -1}}).SetLimit(6)

	jobs, err := fetchJobs(ctx, bson.M{}, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	// Return the list of latest jobs as JSON
	return c.JSON(jobs)
}
