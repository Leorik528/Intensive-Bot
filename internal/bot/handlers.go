package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func HandleMessage(api *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	if msg.Text == "/start" {
		m := tgbotapi.NewMessage(msg.Chat.ID, "Выберите интенсив из списка")
		m.ReplyMarkup = BuildIntensiveKeyboard()
		_, err := api.Send(m)
		return err
	}
	return nil
}
