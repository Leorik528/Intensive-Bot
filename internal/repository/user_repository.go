package repository

import "intensive-bot/internal/domain"

type UserRepository interface {
	CreateOrUpdateTelegramUser(user domain.User) (domain.User, error)
}
