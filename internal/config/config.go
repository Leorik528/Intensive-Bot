package config

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramToken string
	DatabaseURL   string
}

func Load() (*Config, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		// Backward compatibility: allow BOT_TOKEN as an alias.
		telegramToken = os.Getenv("BOT_TOKEN")
	}

	cfg := &Config{
		TelegramToken: telegramToken,
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}

	if cfg.TelegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN (or BOT_TOKEN) is required")
	}
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}
