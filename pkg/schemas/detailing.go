package schemas

import "time"

type GetFinanceDetailingRequest struct {
	DateFrom time.Time `json:"dateFrom"`
	DateTo   time.Time `json:"dateTo"`
}

type ResponseFinanceDetailing struct {
	DateFrom         time.Time `json:"dateFrom"`
	DateTo           time.Time `json:"dateTo"`
	InitialAmount    float64   `json:"initialAmount"`
	TotalIncome      float64   `json:"totalIncome"`
	TotalExpense     float64   `json:"totalExpense"`
	ProfitEstimated  float64   `json:"profitEstimated"`
	ProfitReal       float64   `json:"profitReal"`
	AfterAmountNet   float64   `json:"afterAmountNet"`
	AfterAmountGross float64   `json:"afterAmountGross"`
}

func NewResponseFinanceDetailing(
	dateFrom time.Time,
	dateTo time.Time,
	initialAmount float64,
	totalIncome float64,
	totalExpense float64,
	profitReal float64,
) *ResponseFinanceDetailing {
	profitEstimated := totalIncome - totalExpense
	afterAmountGross := initialAmount + totalIncome
	afterAmountNet := afterAmountGross - totalExpense

	return &ResponseFinanceDetailing{
		DateFrom:         dateFrom,
		DateTo:           dateTo,
		InitialAmount:    initialAmount,
		TotalIncome:      totalIncome,
		TotalExpense:     totalExpense,
		ProfitEstimated:  profitEstimated,
		ProfitReal:       profitReal,
		AfterAmountGross: afterAmountGross,
		AfterAmountNet:   afterAmountNet,
	}
}
