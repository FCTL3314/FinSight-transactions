package domain

import "time"

type RecurringTransaction struct {
	ID                 uint      `gorm:"primaryKey"`
	TransactionID      uint      `gorm:"not null"`
	RecurrenceInterval string    `gorm:"type:interval;not null"`
	IsActive           bool      `gorm:"not null;default:true"`
	CreatedAt          time.Time `gorm:"not null;default:now()"`
	UpdatedAt          time.Time `gorm:"not null;default:now()"`

	Transaction *Transaction `gorm:"foreignKey:TransactionID"`
}
