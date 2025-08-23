package domain

import "time"

type FinanceDetailing struct {
	DateFrom      time.Time
	DateTo        time.Time
	InitialAmount float64
	TotalIncome   float64
	TotalExpense  float64
	ProfitReal    float64
}

func NewFinanceDetailing(
	dateFrom time.Time,
	dateTo time.Time,
	initialAmount float64,
	totalIncome float64,
	totalExpense float64,
	profitReal float64,
) *FinanceDetailing {
	return &FinanceDetailing{
		DateFrom:      dateFrom,
		DateTo:        dateTo,
		InitialAmount: initialAmount,
		TotalIncome:   totalIncome,
		TotalExpense:  totalExpense,
		ProfitReal:    profitReal,
	}
}
