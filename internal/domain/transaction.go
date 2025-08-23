package domain

import (
	"time"
)

type Transaction struct {
	ID         uint
	Amount     float64
	Name       string
	Note       string
	CategoryID int64
	UserID     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time

	RecurringTransaction *RecurringTransaction
}
