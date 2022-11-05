package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nothinux/karsajobs/pkg/models/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type application struct {
	jobs    *mongodb.JobModel
	counter *int
}

func main() {
	client, err := openDB()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	db := client.Database("karsajobs")
	count := 10

	app := &application{
		jobs: &mongodb.JobModel{
			DB: db,
		},
		counter: &count,
	}

	log.Printf("application running on port %s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), app.routes()); err != nil {
		log.Fatal(err)
	}
}

func openDB() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:27017/?authsource=admin", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASS"), os.Getenv("MONGO_HOST"))))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	return client, nil
}
