package domain

import "time"

type RecurringTransaction struct {
	ID                 uint
	TransactionID      uint
	RecurrenceInterval string
	IsActive           bool
	CreatedAt          time.Time
	UpdatedAt          time.Time

	Transaction *Transaction
}
