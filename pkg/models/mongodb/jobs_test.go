package mongodb

import (
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"

	"github.com/nothinux/karsajobs/pkg/models"
)

func TestGetJobPosts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	jp := []*models.JobPost{
		&models.JobPost{
			ID:          ToObjID("5ffe6904eb39b069de45323c"),
			JobID:       2,
			Company:     "PT. Penguins",
			Role:        "Software Engineer",
			Location:    "Jakarta",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
		&models.JobPost{
			ID:          ToObjID("5ffe6904eb39b069de45323b"),
			JobID:       1,
			Company:     "Nothinux",
			Role:        "DevOps Engineer",
			Location:    "Bandung",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
	}

	db, teardown := newTestDB(t)
	defer teardown()

	m := JobModel{db}

	result, err := m.GetJobPosts(context.Background())
	if err != nil {
		t.Errorf("want %v, got %s", jp, err)
	}

	if !reflect.DeepEqual(result, jp) {
		t.Errorf("want %v, got %v", jp, result)
	}

}

func TestGetJobPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	tests := []struct {
		name      string
		ID        string
		wantJob   *models.JobPost
		wantError error
	}{
		{
			name: "Valid ID",
			ID:   "5ffe6904eb39b069de45323b",
			wantJob: &models.JobPost{
				ID:          ToObjID("5ffe6904eb39b069de45323b"),
				JobID:       1,
				Company:     "Nothinux",
				Role:        "DevOps Engineer",
				Location:    "Bandung",
				Description: "no description",
				Status:      true,
				PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:      "Not Valid ID",
			ID:        "601039ed02a32832b7a3eb0d",
			wantJob:   nil,
			wantError: mongo.ErrNoDocuments,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := JobModel{db}

			result, err := m.GetJobPost(context.Background(), tt.ID)
			if err != tt.wantError {
				t.Errorf("want %v; got %v", tt.wantError, err)
			}

			if !reflect.DeepEqual(result, tt.wantJob) {
				t.Errorf("want %v; got %v", tt.wantJob, result)
			}
		})
	}
}

func TestInsertJobPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	data := &models.JobPost{
		JobID:       3,
		Company:     "Iguana Inc",
		Role:        "Fullstack Engineer",
		Location:    "Bandung",
		Description: "no description",
		Status:      true,
		PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
	}

	db, teardown := newTestDB(t)
	defer teardown()

	m := JobModel{db}

	resultid, err := m.InsertJobpost(context.Background(), data.JobID, data.Company, data.Role, data.Location, data.Description, data.Status, data.PublishedAt)
	if err != nil {
		t.Errorf("got %v", err)
	}

	// check data has been inserted
	col := db.Collection("jobposts")
	var jobpost *models.JobPost

	if err := col.FindOne(context.Background(), bson.M{"_id": resultid}).Decode(&jobpost); err != nil {
		t.Errorf("got %v; want %v", resultid, jobpost.ID)
	}
}

func TestDeleteJobPost(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	tests := []struct {
		name             string
		wantID           string
		wantDeletedCount int64
	}{
		{
			name:             "Valid ID",
			wantID:           "5ffe6904eb39b069de45323b",
			wantDeletedCount: 1,
		},
		{
			name:             "Not Valid ID",
			wantID:           "601039ed02a32832b7a3eb0d",
			wantDeletedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := JobModel{db}

			deletedcount, err := m.DeleteJobPost(context.Background(), tt.wantID)
			if err != nil {
				t.Errorf("got %v", err)
			}

			if deletedcount != tt.wantDeletedCount {
				t.Errorf("got %v; want %v", deletedcount, tt.wantDeletedCount)
			}
		})
	}
}
