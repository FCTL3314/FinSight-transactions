package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
)

type DetailingUsecase interface {
	Get(authUserId int64) (*domain.FinanceDetailing, error)
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

func (du *detailingUsecase) Get(authUserId int64) (*domain.FinanceDetailing, error) {
	filterParams := domain.NewFilterParams(
		domain.FilterCondition{
			Field:    "user_id",
			Operator: domain.OpEq,
			Value:    authUserId,
		},
	)

	return du.transactionRepository.GetFinanceDetailing(authUserId, filterParams)
}
