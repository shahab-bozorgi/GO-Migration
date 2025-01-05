package main

import (
    "go-migration-app/config"
    "go-migration-app/internal/api"
    "go-migration-app/internal/db"
    "go-migration-app/internal/middleware"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL == "" {
        log.Fatal("DATABASE_URL is not set in the environment")
    }

    database, err := db.InitDatabase(databaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    server := api.NewServer(config.Config{})

    server.Router.Use(middleware.DatabaseMiddleware(database))
    api.RegisterRoutes(server.Router)

    server.Run(":8080")
}
