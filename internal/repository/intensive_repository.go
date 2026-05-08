package repository

import "intensive-bot/internal/domain"

type IntensiveRepository interface {
	ListOpen() ([]domain.Intensive, error)
	GetByID(id int64) (domain.Intensive, error)
}
