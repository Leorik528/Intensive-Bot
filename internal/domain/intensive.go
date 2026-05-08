package domain

import "time"

type Intensive struct {
	ID          int64
	Title       string
	Description string
	PriceKopeck int64
	StartsAt    time.Time
	ChatID      int64
	IsOpen      bool
}
