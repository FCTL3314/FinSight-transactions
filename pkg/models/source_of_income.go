package models

import "time"

type SourceOfIncome struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:text;not null"`
	Datetime  time.Time `json:"datetime" gorm:"not null"`
	Total     float64   `gorm:"type:decimal(10,2)"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:now()"`
}
