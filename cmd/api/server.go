package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inuoshios/little-jira/internal/config"
	"github.com/inuoshios/little-jira/internal/database"
)

type Application struct {
	config *config.Config
}

func Server(app *Application) error {
	dsn := os.Getenv("DSN")
	conn, err := database.ConnectToDatabase(dsn)
	if err != nil {
		log.Fatalf("error connecting to DB %s", err.Error())
	}
	log.Println("database connected successfully...")

	app.config = &config.Config{
		DB: conn,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: app.routes(),
	}

	idleConnectionClosed := make(chan struct{})
	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		signal.Notify(sigInt, syscall.SIGTERM)
		<-sigInt

		log.Println("service interrupt received")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error %v", err)
		}

		log.Println("shutting down...")
		conn.Close()
		close(idleConnectionClosed)
	}()

	log.Printf("server started at port %s", os.Getenv("PORT"))
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err.Error())
	}

	<-idleConnectionClosed
	log.Println("server shutdown successful")

	return err
}
