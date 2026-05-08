package service

import "github.com/go-co-op/gocron"

type NotificationService struct {
	scheduler *gocron.Scheduler
}

func NewNotificationService(s *gocron.Scheduler) *NotificationService {
	return &NotificationService{scheduler: s}
}
