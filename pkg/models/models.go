package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobPost struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	JobID       int                `bson:"job_id,omitempty"`
	Company     string             `bson:"company,omitempty"`
	Role        string             `bson:"role,omitempty"`
	Location    string             `bson:"location,omitempty"`
	Description string             `bson:"description,omitempty"`
	Status      bool               `bson:"status,omitempty"`
	PublishedAt time.Time          `bson:"published_at,omitempty"`
}

type Health struct {
	Status bool
}
