package models

import "time"

type Transaction struct {
	id         uint64
	Amount     float64
	Name       string
	Note       string
	CategoryID string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type RecurringTransaction struct {
	id            uint64
	TransactionID uint64
	Interval      time.Duration
	IsActive      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PeriodFinancialSummary struct {
	id                   uint64
	DateFrom             time.Time
	DateTo               time.Time
	AmountBefore         float64
	AmountAfterEstimated float64
	AmountAfterFact      float64
	ProfitEstimated      float64
	ProfitFact           float64
}
