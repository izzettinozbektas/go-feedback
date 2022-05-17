package main

import (
	"github.com/izzettinozbektas/go-feedback/cmd/routes"
	"github.com/izzettinozbektas/go-feedback/internal/driver"
	"github.com/izzettinozbektas/go-feedback/internal/handlers"
	"github.com/izzettinozbektas/go-feedback/internal/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

func main() {
	_, err := run()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Handler:      routes.Routes(),
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	helpers.WaitForShutdown(srv)
}

func run() (*mongo.Database, error) {
	db, err := driver.Connect()
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)
	return db, nil
}
