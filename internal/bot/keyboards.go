package bot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"intensive-bot/internal/domain"
)

func BuildIntensiveKeyboard(intensives []domain.Intensive) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, len(intensives))
	for _, in := range intensives {
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s — %d₽", in.Title, in.PriceKopeck/100), "intensive:"+strconv.FormatInt(in.ID, 10)),
		))
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
