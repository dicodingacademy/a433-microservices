package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/nothinux/karsajobs/pkg/models"
	"github.com/nothinux/karsajobs/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newTestDB(t *testing.T) (*application, func()) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017/?authsource=admin", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"), os.Getenv("MONGO_HOST"))))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	// add sample data
	db := client.Database("karsajobs_tests")
	col := db.Collection("jobposts")

	_, err = col.InsertMany(ctx, []interface{}{
		&models.JobPost{
			ID:          mongodb.ToObjID("5fff0b9175eec7e8b33a0611"),
			JobID:       1,
			Company:     "Nothinux",
			Role:        "DevOps Engineer",
			Location:    "Bandung",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
		&models.JobPost{
			ID:          mongodb.ToObjID("5fff0b9175eec7e8b33a0612"),
			JobID:       2,
			Company:     "PT. Penguins",
			Role:        "Software Engineer",
			Location:    "Jakarta",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		jobs: &mongodb.JobModel{
			DB: db,
		},
	}

	return app, func() {
		if err := col.Drop(ctx); err != nil {
			log.Fatal(err)
		}

		client.Disconnect(ctx)
	}

}
