package domain

import "github.com/dromara/carbon/v2"

type FinanceDetailing struct {
	ID               uint
	UserID           int64
	DateFrom         carbon.Carbon
	DateTo           carbon.Carbon
	InitialAmount    float64
	CurrentAmount    float64
	TotalIncome      float64
	TotalExpense     float64
	ProfitEstimated  float64
	ProfitReal       float64
	AfterAmountNet   float64
	AfterAmountGross float64
}

func NewFinanceDetailing(
	dateFrom carbon.Carbon,
	dateTo carbon.Carbon,
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

func (fd *FinanceDetailing) Calculate() {
	fd.ProfitEstimated = fd.TotalIncome - fd.TotalExpense
	fd.ProfitReal = fd.CurrentAmount - fd.InitialAmount
	fd.AfterAmountGross = fd.InitialAmount + fd.TotalIncome
	fd.AfterAmountNet = fd.AfterAmountGross - fd.TotalExpense
}
