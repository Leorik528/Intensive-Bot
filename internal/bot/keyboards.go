package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func BuildIntensiveKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Интенсив #1", "intensive:1"),
		),
	)
}
