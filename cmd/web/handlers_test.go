package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/nothinux/karsajobs/pkg/models/mongodb"

	"github.com/nothinux/karsajobs/pkg/models"
)

func TestHomeHandler(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(home).ServeHTTP(rr, r)
	result := rr.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("got %v, want %v", result.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(string(body), "Up!") {
		t.Errorf("got %v, want %v", string(body), "Up!")
	}

}

func TestGetJobsHandler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	data := []models.JobPost{
		models.JobPost{
			ID:          mongodb.ToObjID("5fff0b9175eec7e8b33a0612"),
			JobID:       2,
			Company:     "PT. Penguins",
			Role:        "Software Engineer",
			Location:    "Jakarta",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
		models.JobPost{
			ID:          mongodb.ToObjID("5fff0b9175eec7e8b33a0611"),
			JobID:       1,
			Company:     "Nothinux",
			Role:        "DevOps Engineer",
			Location:    "Bandung",
			Description: "no description",
			Status:      true,
			PublishedAt: time.Date(2021, 01, 05, 00, 00, 00, 0, time.UTC),
		},
	}

	app, teardown := newTestDB(t)
	defer teardown()

	r, err := http.NewRequest(http.MethodGet, "/jobs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(app.getJobs).ServeHTTP(rr, r)
	result := rr.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("got %v, want %v", result.StatusCode, http.StatusOK)
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	var jp []models.JobPost

	if err := json.Unmarshal(body, &jp); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(jp, data) {
		t.Errorf("got %v, want %v", jp, data)
	}
}
