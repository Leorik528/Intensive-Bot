package service

import (
	"intensive-bot/internal/domain"
	"intensive-bot/internal/repository"
)

type IntensiveService struct {
	repo repository.IntensiveRepository
}

func NewIntensiveService(repo repository.IntensiveRepository) *IntensiveService {
	return &IntensiveService{repo: repo}
}

func (s *IntensiveService) ListOpen() ([]domain.Intensive, error) {
	return s.repo.ListOpen()
}
