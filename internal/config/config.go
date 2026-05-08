package config

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramToken string
	ProviderToken string
	DatabaseURL   string
	AdminHTTPAddr string
}

func Load() (*Config, error) {
	cfg := &Config{TelegramToken: os.Getenv("TELEGRAM_TOKEN"), ProviderToken: os.Getenv("PAYMENT_PROVIDER_TOKEN"), DatabaseURL: os.Getenv("DATABASE_URL"), AdminHTTPAddr: os.Getenv("ADMIN_HTTP_ADDR")}
	if cfg.AdminHTTPAddr == "" {
		cfg.AdminHTTPAddr = ":8080"
	}
	if cfg.TelegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN is required")
	}
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.ProviderToken == "" {
		return nil, fmt.Errorf("PAYMENT_PROVIDER_TOKEN is required")
	}
	return cfg, nil
}
