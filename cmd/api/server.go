package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/inuoshios/little-jira/internal/database"
	"github.com/joho/godotenv"
)

type Application struct {
	DB *sql.DB
}

func Server() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	conn, err := database.ConnectToDatabase(dsn)
	if err != nil {
		log.Fatalf("Eror connecting to DB %s", err.Error())
	}

	log.Println("Database connected successfully...")

	return conn, nil
}
