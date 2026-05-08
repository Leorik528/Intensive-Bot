package service

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"intensive-bot/internal/domain"
)

type PaymentService struct {
	providerToken string
}

func NewPaymentService(providerToken string) *PaymentService {
	return &PaymentService{providerToken: providerToken}
}

func (s *PaymentService) BuildInvoice(chatID int64, intensive domain.Intensive) tgbotapi.InvoiceConfig {
	price := tgbotapi.LabeledPrice{Label: intensive.Title, Amount: int(intensive.PriceKopeck)}
	return tgbotapi.NewInvoice(chatID, intensive.Title, intensive.Description, strconv.FormatInt(intensive.ID, 10), s.providerToken, "RUB", price)
}

func (s *PaymentService) ValidatePreCheckout(q *tgbotapi.PreCheckoutQuery) error {
	if q == nil {
		return fmt.Errorf("empty pre-checkout query")
	}
	if q.InvoicePayload == "" {
		return fmt.Errorf("empty invoice payload")
	}
	return nil
}

func (s *PaymentService) ParsePaidIntensiveID(payload string) (int64, error) {
	return strconv.ParseInt(payload, 10, 64)
}
