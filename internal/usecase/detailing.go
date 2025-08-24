package usecase

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
)

type DetailingUsecase interface {
	Get(authUserId int64, dateFrom, dateTo time.Time) (*domain.FinanceDetailing, error)
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

func (du *detailingUsecase) Get(authUserId int64, dateFrom, dateTo time.Time) (*domain.FinanceDetailing, error) {
	filterParams := domain.NewFilterParams(
		domain.FilterCondition{
			Field:    "user_id",
			Operator: domain.OpEq,
			Value:    authUserId,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpGte,
			Value:    dateFrom,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpLte,
			Value:    dateTo,
		},
	)

	return du.transactionRepository.GetFinanceDetailing(dateFrom, dateTo, filterParams)
}
