package models

import "time"

type PeriodFinancialSummary struct {
	ID                 uint      `json:"id" gorm:"primaryKey"`
	DateFrom           time.Time `json:"date_from" gorm:"type:date;not null"`
	DateTo             time.Time `json:"date_to" gorm:"type:date;not null"`
	StartingBalance    float64   `json:"starting_balance" gorm:"type:numeric(12,2);not null"`
	ProjectedBalance   float64   `json:"projected_balance" gorm:"type:numeric(12,2)"`
	ActualBalance      float64   `json:"actual_balance" gorm:"type:numeric(12,2)"`
	ProjectedNetChange float64   `json:"projected_net_change" gorm:"type:numeric(12,2)"`
	ActualNetChange    float64   `json:"actual_net_change" gorm:"type:numeric(12,2)"`
	CreatedAt          time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"not null;default:now()"`
}
