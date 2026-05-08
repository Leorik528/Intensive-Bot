package app

import (
	"context"

	"intensive-bot/internal/bot"
	"intensive-bot/internal/config"
)

type App struct {
	cfg *config.Config
	bot *bot.Bot
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	b, err := bot.New(cfg)
	if err != nil {
		return nil, err
	}

	return &App{cfg: cfg, bot: b}, nil
}

func (a *App) Run() error {
	return a.bot.Run(context.Background())
}
