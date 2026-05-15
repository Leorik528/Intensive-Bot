package bot

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"intensive-bot/internal/config"
	"intensive-bot/internal/service"
)

type Bot struct {
	api             *tgbotapi.BotAPI
	yooKassaService *service.YooKassaService
}

func New(cfg *config.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	yooKassa := service.NewYooKassaService(cfg.YooKassaShopID, cfg.YooKassaSecretKey, cfg.YooKassaReturnURL)

	return &Bot{api: api, yooKassaService: yooKassa}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	_ = ctx
	updates := b.api.GetUpdatesChan(tgbotapi.NewUpdate(0))
	for update := range updates {
		if update.Message != nil {
			_ = HandleMessage(b.api, update.Message)
			continue
		}

		if update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, "intensive:") {
			if err := HandleIntensiveSelection(b.api, b.yooKassaService, update.CallbackQuery); err != nil {
				_ = HandleErrorMessage(b.api, update.CallbackQuery.Message.Chat.ID, "Не удалось создать платеж. Попробуйте позже")
			}
		}
	}
	return nil
}
