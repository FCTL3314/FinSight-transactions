package schemas

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/dromara/carbon/v2"
)

type GetFinanceDetailingRequest struct {
	DateFrom      carbon.Carbon `form:"date_from" binding:"required" time_format:"2006-01-02"`
	DateTo        carbon.Carbon `form:"date_to" binding:"required" time_format:"2006-01-02"`
	InitialAmount float64       `form:"initial_amount" binding:"required"`
	CurrentAmount float64       `form:"current_amount" binding:"required"`
}

type CreateFinanceDetailingRequest struct {
	DateFrom      carbon.Carbon `json:"date_from" binding:"required,datetime=2006-01-02"`
	DateTo        carbon.Carbon `json:"date_to" binding:"required,datetime=2006-01-02"`
	InitialAmount float64       `json:"initial_amount" binding:"required"`
	CurrentAmount float64       `json:"current_amount" binding:"required"`
}

func (req *CreateFinanceDetailingRequest) ToDomainModel(userID int64) *domain.FinanceDetailing {
	return &domain.FinanceDetailing{
		UserID:        userID,
		DateFrom:      req.DateFrom,
		DateTo:        req.DateTo,
		InitialAmount: req.InitialAmount,
		CurrentAmount: req.CurrentAmount,
	}
}

type UpdateFinanceDetailingRequest struct {
	DateFrom      *carbon.Carbon `json:"date_from" binding:"omitempty,datetime=2006-01-02"`
	DateTo        *carbon.Carbon `json:"date_to" binding:"omitempty,datetime=2006-01-02"`
	InitialAmount *float64       `json:"initial_amount"`
	CurrentAmount *float64       `json:"current_amount"`
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
	ID               uint          `json:"id"`
	DateFrom         carbon.Carbon `json:"date_from" time_format:"2006-01-02"`
	DateTo           carbon.Carbon `json:"date_to" time_format:"2006-01-02"`
	InitialAmount    float64       `json:"initial_amount"`
	CurrentAmount    float64       `json:"current_amount"`
	TotalIncome      float64       `json:"total_income"`
	TotalExpense     float64       `json:"total_expense"`
	ProfitEstimated  float64       `json:"profit_estimated"`
	ProfitReal       float64       `json:"profit_real"`
	AfterAmountNet   float64       `json:"after_amount_net"`
	AfterAmountGross float64       `json:"after_amount_gross"`
}

func NewResponseFinanceDetailing(fd *domain.FinanceDetailing) *ResponseFinanceDetailing {
	if fd == nil {
		return nil
	}

	profitReal := fd.CurrentAmount - fd.InitialAmount
	afterAmountGross := fd.InitialAmount + fd.TotalIncome
	afterAmountNet := afterAmountGross - fd.TotalExpense

	return &ResponseFinanceDetailing{
		ID:               fd.ID,
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

func NewResponseFinanceDetailingList(detailings []*domain.FinanceDetailing) []*ResponseFinanceDetailing {
	response := make([]*ResponseFinanceDetailing, len(detailings))
	for i, detailing := range detailings {
		response[i] = NewResponseFinanceDetailing(detailing)
	}
	return response
}
