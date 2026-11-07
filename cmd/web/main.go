package main

import (
	"log"
	"net/http"

	"github.com/fazilazil184-dev/GoPanel/internal/config"
	"github.com/fazilazil184-dev/GoPanel/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}

	db := config.ConnectDB()

	r := routes.SetupRouter(db)

	log.Printf("App running on port: %v\n", 2026)
	log.Fatal(http.ListenAndServe(":2026", r))
}
