package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/schemas"
)

type DetailingUsecase interface {
	Get(authUserId int64, getFinanceDetailingRequest *schemas.GetFinanceDetailingRequest) (*domain.FinanceDetailing, error)
}

type detailingUsecase struct {
	transactionRepository repository.TransactionRepository
	cfg                   *config.Config
}

func NewDetailingUsecase(
	transactionRepository repository.TransactionRepository,
	cfg *config.Config,
) DetailingUsecase {
	return &detailingUsecase{
		transactionRepository: transactionRepository,
		cfg:                   cfg,
	}
}

func (du *detailingUsecase) Get(authUserId int64, getFinanceDetailingRequest *schemas.GetFinanceDetailingRequest) (*domain.FinanceDetailing, error) {
	filterParams := domain.NewFilterParams(
		domain.FilterCondition{
			Field:    "user_id",
			Operator: domain.OpEq,
			Value:    authUserId,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpGte,
			Value:    getFinanceDetailingRequest.DateFrom,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpLte,
			Value:    getFinanceDetailingRequest.DateTo,
		},
	)

	return du.transactionRepository.GetFinanceDetailing(
		getFinanceDetailingRequest.DateFrom,
		getFinanceDetailingRequest.DateTo,
		getFinanceDetailingRequest.InitialAmount,
		getFinanceDetailingRequest.CurrentAmount,
		filterParams,
	)
}
