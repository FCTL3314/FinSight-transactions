package schemas

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
)

type CreateFinanceDetailingRequest struct {
	DateFrom      time.Time `json:"date_from" time_format:"2006-01-02" binding:"required"`
	DateTo        time.Time `json:"date_to" time_format:"2006-01-02" binding:"required"`
	InitialAmount float64   `json:"initial_amount" binding:"required"`
	CurrentAmount float64   `json:"current_amount" binding:"required"`
}

func (req *CreateFinanceDetailingRequest) ToDomainModel(userID int64) *domain.FinanceDetailing {
	return domain.NewFinanceDetailing(
		userID,
		req.DateFrom,
		req.DateTo,
		req.InitialAmount,
		req.CurrentAmount,
		0,
		0,
	)
}

type UpdateFinanceDetailingRequest struct {
	DateFrom      *time.Time `json:"date_from" time_format:"2006-01-02" binding:"omitempty"`
	DateTo        *time.Time `json:"date_to" time_format:"2006-01-02" binding:"omitempty"`
	InitialAmount *float64   `json:"initial_amount"`
	CurrentAmount *float64   `json:"current_amount"`
}

func (req *UpdateFinanceDetailingRequest) ApplyToDomainModel(fd *domain.FinanceDetailing) {
	if req.DateFrom != nil {
		fd.DateFrom = *req.DateFrom
	}
	if req.DateTo != nil {
		fd.DateTo = *req.DateTo
	}
	if req.InitialAmount != nil {
		fd.InitialAmount = *req.InitialAmount
	}
	if req.CurrentAmount != nil {
		fd.CurrentAmount = *req.CurrentAmount
	}
}

type ResponseFinanceDetailing struct {
	ID               uint      `json:"id"`
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

	fd.Calculate()

	return &ResponseFinanceDetailing{
		ID:               fd.ID,
		DateFrom:         fd.DateFrom,
		DateTo:           fd.DateTo,
		InitialAmount:    fd.InitialAmount,
		CurrentAmount:    fd.CurrentAmount,
		TotalIncome:      fd.TotalIncome,
		TotalExpense:     fd.TotalExpense,
		ProfitEstimated:  fd.ProfitEstimated,
		ProfitReal:       fd.ProfitReal,
		AfterAmountNet:   fd.AfterAmountNet,
		AfterAmountGross: fd.AfterAmountGross,
	}
}

func NewResponseFinanceDetailingList(detailings []*domain.FinanceDetailing) []*ResponseFinanceDetailing {
	response := make([]*ResponseFinanceDetailing, len(detailings))
	for i, detailing := range detailings {
		response[i] = NewResponseFinanceDetailing(detailing)
	}
	return response
}
