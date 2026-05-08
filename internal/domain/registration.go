package domain

import "time"

type Registration struct {
	ID          int64
	UserID      int64
	IntensiveID int64
	PaidAt      *time.Time
	AccessLink  string
	Status      string
}
