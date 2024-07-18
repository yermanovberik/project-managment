package main

import (
	"log"
	"net/http"
	"project-managment/internal/app/config"
	"project-managment/internal/app/db"
	"project-managment/internal/app/handlers"
	"project-managment/internal/app/repository"
	"project-managment/internal/app/routes"
)

func main() {
	cfg := config.LoadConfig()
	database, err := db.InitDB(cfg)

	if err != nil {
		log.Fatal("failed to init database")
	}

	repo := repository.NewRepository(database)

	handler := handlers.NewHandler(repo)

	router := routes.SetupRoutes(handler)

	log.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
