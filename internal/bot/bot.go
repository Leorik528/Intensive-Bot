package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"intensive-bot/internal/config"
	"intensive-bot/internal/domain"
	"intensive-bot/internal/service"
)

type Bot struct {
	api              *tgbotapi.BotAPI
	intensiveService *service.IntensiveService
	paymentService   *service.PaymentService
	accessService    *service.AccessService
}

func New(cfg *config.Config, intensiveService *service.IntensiveService) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}
	return &Bot{api: api, intensiveService: intensiveService, paymentService: service.NewPaymentService(cfg.ProviderToken), accessService: service.NewAccessService()}, nil
}

func (b *Bot) Run(ctx context.Context) error {
	_ = ctx
	updates := b.api.GetUpdatesChan(tgbotapi.NewUpdate(0))
	for update := range updates {
		switch {
		case update.Message != nil:
			_ = b.HandleMessage(update.Message)
		case update.CallbackQuery != nil:
			_ = b.HandleCallback(update.CallbackQuery)
		case update.PreCheckoutQuery != nil:
			_ = b.HandlePreCheckout(update.PreCheckoutQuery)
		}
	}
	return nil
}

func (b *Bot) seed() {
	if len(b.intensiveService.ListAll()) > 0 {
		return
	}
	b.intensiveService.Create(domain.Intensive{Title: "Go Intensive", Description: "MVP поток", PriceKopeck: 199000, ChatID: -1001234567890, IsOpen: true})
}

func (b *Bot) HandleMessage(msg *tgbotapi.Message) error {
	b.seed()
	if msg.SuccessfulPayment != nil {
		id, err := b.paymentService.ParsePaidIntensiveID(msg.SuccessfulPayment.InvoicePayload)
		if err != nil {
			return err
		}
		intensive, err := b.intensiveService.GetByID(id)
		if err != nil {
			return err
		}
		link, err := b.accessService.CreateInviteLink(b.api, intensive.ChatID, msg.From.ID)
		if err != nil {
			return err
		}
		_, err = b.api.Send(tgbotapi.NewMessage(msg.Chat.ID, "Оплата получена. Ссылка в закрытый чат: "+link))
		return err
	}
	if msg.Text == "/start" {
		list, _ := b.intensiveService.ListOpen()
		m := tgbotapi.NewMessage(msg.Chat.ID, "Выберите интенсив из списка")
		m.ReplyMarkup = BuildIntensiveKeyboard(list)
		_, err := b.api.Send(m)
		return err
	}
	if strings.HasPrefix(msg.Text, "/refund ") {
		parts := strings.Split(msg.Text, " ")
		if len(parts) == 3 {
			chatID, _ := strconv.ParseInt(parts[1], 10, 64)
			userID, _ := strconv.ParseInt(parts[2], 10, 64)
			return b.accessService.BanMember(b.api, chatID, userID)
		}
	}
	return nil
}

func (b *Bot) HandleCallback(q *tgbotapi.CallbackQuery) error {
	if strings.HasPrefix(q.Data, "intensive:") {
		id, _ := strconv.ParseInt(strings.TrimPrefix(q.Data, "intensive:"), 10, 64)
		intensive, err := b.intensiveService.GetByID(id)
		if err != nil {
			return err
		}
		invoice := b.paymentService.BuildInvoice(q.Message.Chat.ID, intensive)
		if _, err := b.api.Send(invoice); err != nil {
			return err
		}
		_, _ = b.api.Request(tgbotapi.NewCallback(q.ID, fmt.Sprintf("Выбран: %s", intensive.Title)))
		return nil
	}
	if strings.HasPrefix(q.Data, "regen:") {
		id, _ := strconv.ParseInt(strings.TrimPrefix(q.Data, "regen:"), 10, 64)
		intensive, err := b.intensiveService.GetByID(id)
		if err != nil {
			return err
		}
		link, err := b.accessService.CreateInviteLink(b.api, intensive.ChatID, q.From.ID)
		if err != nil {
			return err
		}
		_, err = b.api.Send(tgbotapi.NewMessage(q.Message.Chat.ID, "Новая ссылка: "+link))
		return err
	}
	return nil
}

func (b *Bot) HandlePreCheckout(q *tgbotapi.PreCheckoutQuery) error {
	err := b.paymentService.ValidatePreCheckout(q)
	cfg := tgbotapi.PreCheckoutConfig{PreCheckoutQueryID: q.ID, OK: err == nil}
	if err != nil {
		cfg.ErrorMessage = err.Error()
	}
	_, reqErr := b.api.Request(cfg)
	if reqErr != nil {
		return reqErr
	}
	return err
}
