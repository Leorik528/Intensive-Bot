package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const yookassaBaseURL = "https://api.yookassa.ru/v3/payments"

type YooKassaService struct {
	shopID    string
	secretKey string
	returnURL string
	http      *http.Client
}

func NewYooKassaService(shopID, secretKey, returnURL string) *YooKassaService {
	return &YooKassaService{
		shopID:    shopID,
		secretKey: secretKey,
		returnURL: returnURL,
		http:      &http.Client{Timeout: 15 * time.Second},
	}
}

func (s *YooKassaService) CreatePayment(telegramUserID int64, amountRub int64, description string) (string, error) {
	if amountRub <= 0 {
		return "", fmt.Errorf("amount must be positive")
	}

	payload := map[string]any{
		"amount": map[string]string{
			"value":    fmt.Sprintf("%d.00", amountRub),
			"currency": "RUB",
		},
		"capture":     true,
		"description": description,
		"metadata": map[string]string{
			"telegram_user_id": strconv.FormatInt(telegramUserID, 10),
		},
		"confirmation": map[string]string{
			"type":       "redirect",
			"return_url": s.returnURL,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, yookassaBaseURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(s.shopID, s.secretKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotence-Key", fmt.Sprintf("tg-%d-%d", telegramUserID, time.Now().UnixNano()))

	resp, err := s.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("yookassa create payment failed: status %d", resp.StatusCode)
	}

	var out struct {
		Confirmation struct {
			URL string `json:"confirmation_url"`
		} `json:"confirmation"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if out.Confirmation.URL == "" {
		return "", fmt.Errorf("empty confirmation_url in yookassa response")
	}

	return out.Confirmation.URL, nil
}
