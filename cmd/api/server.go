package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/inuoshios/little-jira/internal/config"
	"github.com/inuoshios/little-jira/internal/database"
)

type Application struct {
	config *config.Config
}

func Server() error {
	dsn := os.Getenv("DSN")
	conn, err := database.ConnectToDatabase(dsn)
	if err != nil {
		log.Fatalf("error connecting to DB %s", err.Error())
	}
	log.Println("database connected successfully...")

	app := Application{
		config: &config.Config{
			DB: conn,
		},
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}
	log.Println("server running successfully")

	return nil
}
