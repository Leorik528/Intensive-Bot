package repository

import "intensive-bot/internal/domain"

type RegistrationRepository interface {
	CreatePaidRegistration(reg domain.Registration) (domain.Registration, error)
	ListPaidByIntensive(intensiveID int64) ([]domain.Registration, error)
}
