package config

import (
	"fmt"
	"os"
)

type Config struct {
	TelegramToken     string
	DatabaseURL       string
	YooKassaShopID    string
	YooKassaSecretKey string
	YooKassaReturnURL string
}

func Load() (*Config, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	if telegramToken == "" {
		// Backward compatibility: allow BOT_TOKEN as an alias.
		telegramToken = os.Getenv("BOT_TOKEN")
	}

	cfg := &Config{
		TelegramToken:     telegramToken,
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		YooKassaShopID:    os.Getenv("YOOKASSA_SHOP_ID"),
		YooKassaSecretKey: os.Getenv("YOOKASSA_SECRET_KEY"),
		YooKassaReturnURL: os.Getenv("YOOKASSA_RETURN_URL"),
	}

	if cfg.TelegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN (or BOT_TOKEN) is required")
	}
	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.YooKassaShopID == "" {
		return nil, fmt.Errorf("YOOKASSA_SHOP_ID is required")
	}
	if cfg.YooKassaSecretKey == "" {
		return nil, fmt.Errorf("YOOKASSA_SECRET_KEY is required")
	}

	return cfg, nil
}
