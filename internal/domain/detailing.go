package domain

import "time"

type FinanceDetailing struct {
	DateFrom        time.Time
	DateTo          time.Time
	InitialAmount   float64
	CurrentAmount   float64
	TotalIncome     float64
	TotalExpense    float64
	ProfitEstimated float64
}

func NewFinanceDetailing(
	dateFrom time.Time,
	dateTo time.Time,
	initialAmount float64,
	currentAmount float64,
	totalIncome float64,
	totalExpense float64,
	profitEstimated float64,
) *FinanceDetailing {
	return &FinanceDetailing{
		DateFrom:        dateFrom,
		DateTo:          dateTo,
		InitialAmount:   initialAmount,
		CurrentAmount:   currentAmount,
		TotalIncome:     totalIncome,
		TotalExpense:    totalExpense,
		ProfitEstimated: profitEstimated,
	}
}
