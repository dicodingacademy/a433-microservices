package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nothinux/karsajobs/pkg/models"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type JobModel struct {
	DB *mongo.Database
}

func (m *JobModel) GetJobPosts(ctx context.Context) ([]*models.JobPost, error) {
	col := m.DB.Collection("jobposts")

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})

	data, err := col.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	var jobposts []*models.JobPost

	if err = data.All(ctx, &jobposts); err != nil {
		return nil, err
	}

	return jobposts, nil
}

func (m *JobModel) GetJobPost(ctx context.Context, id string) (*models.JobPost, error) {
	col := m.DB.Collection("jobposts")
	var jobpost *models.JobPost

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := col.FindOne(ctx, bson.M{"_id": objectid}).Decode(&jobpost); err != nil {
		return nil, err
	}

	return jobpost, nil
}

func (m *JobModel) InsertJobpost(ctx context.Context, id int, company, role, location,
	description string, status bool, publishedat time.Time) (interface{}, error) {
	col := m.DB.Collection("jobposts")

	jobpost := &models.JobPost{
		JobID:       id,
		Company:     company,
		Role:        role,
		Location:    location,
		Description: description,
		Status:      status,
		PublishedAt: publishedat,
	}

	result, err := col.InsertOne(ctx, jobpost)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (m *JobModel) DeleteJobPost(ctx context.Context, id string) (int64, error) {
	col := m.DB.Collection("jobposts")

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	result, err := col.DeleteOne(ctx, bson.M{"_id": objectid})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
