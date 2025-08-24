package usecase

import (
	"fmt"
	"time"

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
	params := domain.NewParams(filterParams, nil, nil)

	transactions, err := du.transactionRepository.Fetch(params)
	if err != nil {
		return nil, err
	}

	fmt.Println(transactions)
	// TODO: Здесь должна быть логика расчета детализации на основе полученных транзакций
	return domain.NewFinanceDetailing(time.Time{}, time.Time{}, 0, 0, 0, 0), nil
}
