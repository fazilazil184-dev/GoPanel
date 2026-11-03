package main

import (
	"log"

	"github.com/fazilazil184-dev/GoPanel/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Failed to load .env file")
	}

	config.ConnectDB()

}
