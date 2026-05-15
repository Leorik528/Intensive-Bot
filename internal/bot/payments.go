package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"intensive-bot/internal/service"
)

func HandleIntensiveSelection(api *tgbotapi.BotAPI, yoo *service.YooKassaService, cb *tgbotapi.CallbackQuery) error {
	_, _ = api.Request(tgbotapi.NewCallback(cb.ID, "Готовим ссылку на оплату..."))

	url, err := yoo.CreatePayment(cb.From.ID, 1000, fmt.Sprintf("Оплата интенсива %s", cb.Data))
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(cb.Message.Chat.ID, "Оплатите интенсив по ссылке:\n"+url)
	_, err = api.Send(msg)
	return err
}

func HandleErrorMessage(api *tgbotapi.BotAPI, chatID int64, text string) error {
	_, err := api.Send(tgbotapi.NewMessage(chatID, text))
	return err
}
