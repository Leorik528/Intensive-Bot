package main

import (
	"log"

	"intensive-bot/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env file not loaded: %v", err)
	}

	a, err := app.New()
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
