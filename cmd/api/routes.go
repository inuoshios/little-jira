package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"author":          "Inu John",
			"github_username": "https://github.com/inuoshios",
			"message":         "little jira api...",
		}
		marshallResponse, _ := json.MarshalIndent(response, "", "\t")

		w.WriteHeader(http.StatusOK)
		w.Write(marshallResponse)
	})

	mux.Get("/get-users", app.handlers.GetUsers)

	return mux
}
