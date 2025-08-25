package schemas

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
)

type GetFinanceDetailingRequest struct {
	DateFrom      time.Time `form:"date_from" binding:"required"`
	DateTo        time.Time `form:"date_to" binding:"required"`
	InitialAmount float64   `form:"initial_amount" binding:"required"`
	CurrentAmount float64   `form:"current_amount" binding:"required"`
}

type ResponseFinanceDetailing struct {
	DateFrom         time.Time `json:"date_from"`
	DateTo           time.Time `json:"date_to"`
	InitialAmount    float64   `json:"initial_amount"`
	CurrentAmount    float64   `json:"current_amount"`
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

	profitReal := fd.CurrentAmount - fd.InitialAmount
	afterAmountGross := fd.InitialAmount + fd.TotalIncome
	afterAmountNet := afterAmountGross - fd.TotalExpense

	return &ResponseFinanceDetailing{
		DateFrom:         fd.DateFrom,
		DateTo:           fd.DateTo,
		InitialAmount:    fd.InitialAmount,
		CurrentAmount:    fd.CurrentAmount,
		TotalIncome:      fd.TotalIncome,
		TotalExpense:     fd.TotalExpense,
		ProfitEstimated:  fd.ProfitEstimated,
		ProfitReal:       profitReal,
		AfterAmountGross: afterAmountGross,
		AfterAmountNet:   afterAmountNet,
	}
}
