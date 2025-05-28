package models

import "time"

type Transaction struct {
	ID         uint                 `gorm:"primaryKey" json:"id"`
	Amount     float64              `gorm:"type:numeric(12,2);not null" json:"amount"`
	Name       string               `gorm:"type:text;not null" json:"name"`
	Note       *string              `gorm:"type:text" json:"note,omitempty"`
	CategoryID uint                 `gorm:"not null" json:"category_id"`
	UserID     uint                 `gorm:"not null" json:"user_id"`
	CreatedAt  time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
	Recurring  RecurringTransaction `gorm:"foreignKey:TransactionID" json:"recurring,omitempty"`
}
type RecurringTransaction struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	TransactionID      uint      `gorm:"not null;index" json:"transaction_id"`
	RecurrenceInterval string    `gorm:"type:interval;not null" json:"recurrence_interval"`
	IsActive           bool      `gorm:"not null;default:true" json:"is_active"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type PeriodFinancialSummary struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	DateFrom           time.Time `gorm:"type:date;not null" json:"date_from"`
	DateTo             time.Time `gorm:"type:date;not null" json:"date_to"`
	StartingBalance    float64   `gorm:"type:numeric(12,2);not null" json:"starting_balance"`
	ProjectedBalance   float64   `gorm:"type:numeric(12,2)" json:"projected_balance"`
	ActualBalance      float64   `gorm:"type:numeric(12,2)" json:"actual_balance"`
	ProjectedNetChange float64   `gorm:"type:numeric(12,2)" json:"projected_net_change"`
	ActualNetChange    float64   `gorm:"type:numeric(12,2)" json:"actual_net_change"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
