package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AccessService struct{}

func NewAccessService() *AccessService { return &AccessService{} }

func (s *AccessService) CreateInviteLink(api *tgbotapi.BotAPI, chatID, userID int64) (string, error) {
	cfg := tgbotapi.CreateChatInviteLinkConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatID}, Name: fmt.Sprintf("user_%d", userID), MemberLimit: 1, ExpireDate: time.Now().Add(24 * time.Hour).Unix()}
	link, err := api.Request(cfg)
	if err != nil {
		return "", err
	}
	if !link.Ok {
		return "", fmt.Errorf("telegram api error")
	}
	invite, ok := link.Result.(map[string]any)
	if !ok {
		return "", fmt.Errorf("bad invite payload")
	}
	url, _ := invite["invite_link"].(string)
	return url, nil
}

func (s *AccessService) BanMember(api *tgbotapi.BotAPI, chatID, userID int64) error {
	_, err := api.Request(tgbotapi.BanChatMemberConfig{ChatMemberConfig: tgbotapi.ChatMemberConfig{ChatConfig: tgbotapi.ChatConfig{ChatID: chatID}, UserID: userID}})
	return err
}
