package main

import (
	"log"

	"intensive-bot/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("app init failed: %v", err)
	}

	if err := a.Run(); err != nil {
		log.Fatalf("app run failed: %v", err)
	}
}
