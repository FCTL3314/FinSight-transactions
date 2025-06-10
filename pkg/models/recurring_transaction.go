package models

import "time"

type RecurringTransaction struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	TransactionID      uint      `json:"transaction_id" gorm:"not null"`
	RecurrenceInterval string    `json:"recurrence_interval" gorm:"type:interval;not null"`
	IsActive           bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt          time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"not null;default:now()"`

	Transaction *Transaction `json:"transaction,omitempty" gorm:"foreignKey:TransactionID"`
}
