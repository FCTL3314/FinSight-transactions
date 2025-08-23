package schemas

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/pkg/models"
)

type ResponseRecurringTransaction struct {
	ID                 uint      `json:"id"`
	TransactionID      uint      `json:"transaction_id"`
	RecurrenceInterval string    `json:"recurrence_interval"`
	IsActive           bool      `json:"is_active"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	Transaction *models.Transaction `json:"transaction,omitempty"`
}
