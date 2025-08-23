package schemas

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
)

type GetFinanceDetailingRequest struct {
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

type ResponseFinanceDetailing struct {
	DateFrom         time.Time `json:"date_from"`
	DateTo           time.Time `json:"date_to"`
	InitialAmount    float64   `json:"initial_amount"`
	TotalIncome      float64   `json:"total_income"`
	TotalExpense     float64   `json:"total_expense"`
	ProfitEstimated  float64   `json:"profit_estimated"`
	ProfitReal       float64   `json:"profit_real"`
	AfterAmountNet   float64   `json:"after_amount_net"`
	AfterAmountGross float64   `json:"after_amount_gross"`
}

func NewResponseFinanceDetailing(fd *domain.FinanceDetailing) *ResponseFinanceDetailing {
	if fd == nil {
		return nil
	}

	profitEstimated := fd.TotalIncome - fd.TotalExpense
	afterAmountGross := fd.InitialAmount + fd.TotalIncome
	afterAmountNet := afterAmountGross - fd.TotalExpense

	return &ResponseFinanceDetailing{
		DateFrom:         fd.DateFrom,
		DateTo:           fd.DateTo,
		InitialAmount:    fd.InitialAmount,
		TotalIncome:      fd.TotalIncome,
		TotalExpense:     fd.TotalExpense,
		ProfitEstimated:  profitEstimated,
		ProfitReal:       fd.ProfitReal,
		AfterAmountGross: afterAmountGross,
		AfterAmountNet:   afterAmountNet,
	}
}
