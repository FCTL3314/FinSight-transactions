package schemas

import (
	"time"
)

type ResponseRecurringTransaction struct {
	ID                 uint      `json:"id"`
	TransactionID      uint      `json:"transaction_id"`
	RecurrenceInterval string    `json:"recurrence_interval"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	Transaction *ResponseTransaction `json:"transaction,omitempty"`
}
