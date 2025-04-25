package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobMode string

const (
	Remote JobMode = "remote"
	Onsite JobMode = "onsite"
	Hybrid JobMode = "hybrid"
)

type JobType string

const (
	Internship JobType = "internship"
	PartTime   JobType = "part_time"
	FullTime   JobType = "full_time"
)

type Jobs struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Salary         string             `bson:"salary,omitempty" json:"salary"`
	Username       string             `bson:"username" json:"username"`
	JobRole        string             `bson:"job_role" json:"job_role"`
	CompanyName    string             `bson:"company_name" json:"company_name"`
	Location       string             `bson:"location", json:"location"`
	JobType        JobType            `bson:"job_type" json:"job_type"`
	JobMode        JobMode            `bson:"job_mode" json:"job_mode"`
	Validity       time.Time          `bson:"validity" json:"validity"`
	JobLink        string             `bson:"job_link" json:"job_link"`
	JobDescription string             `bson:"job_description" json:"job_description"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
}
