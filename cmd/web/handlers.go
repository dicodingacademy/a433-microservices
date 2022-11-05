package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/go-chi/chi"
	"github.com/nothinux/karsajobs/pkg/models"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-type", "text/plain")
	w.Write([]byte("Up!"))
}

func (app *application) getJobs(w http.ResponseWriter, r *http.Request) {
	data, err := app.jobs.GetJobPosts(context.Background())
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(out)
}

func (app *application) getJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := app.jobs.GetJobPost(context.Background(), id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	out, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(out)
}

func (app *application) InsertJob(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var jp models.JobPost
	if err := json.Unmarshal(body, &jp); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// increse counter for job_id
	app.increment()

	_, err = app.jobs.InsertJobpost(context.Background(), app.getCount(), jp.Company, jp.Role, jp.Location, jp.Description, jp.Status, time.Now())
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"status": "ok"}`))

}

func (app *application) DeleteJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := app.jobs.DeleteJobPost(context.Background(), id)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")

	res := strconv.Itoa(int(data))
	w.Write([]byte(res))
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	client, err := openDB()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	defer client.Disconnect(ctx)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	out, err := json.Marshal(&models.Health{
		Status: true,
	})

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(out)
}
