package bot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"intensive-bot/internal/config"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func New(cfg *config.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	return &Bot{api: api}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	_ = ctx
	updates := b.api.GetUpdatesChan(tgbotapi.NewUpdate(0))
	for update := range updates {
		if update.Message == nil {
			continue
		}
		_ = HandleMessage(b.api, update.Message)
	}
	return nil
}
