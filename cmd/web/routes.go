package main

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	}))

	mux.Get("/", home)
	mux.Get("/jobs", app.getJobs)
	mux.Get("/job/{id}", app.getJob)
	mux.Post("/job", app.InsertJob)
	mux.Delete("/job/{id}", app.DeleteJob)
	mux.Get("/health", app.health)

	return mux
}
