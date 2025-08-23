package usecase

import (
	"fmt"
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
)

type DetailingUsecase interface {
	Get(authUserId int64, params *domain.Params) (*domain.FinanceDetailing, error)
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

func (du *detailingUsecase) Get(authUserId int64, params *domain.Params) (*domain.FinanceDetailing, error) {
	authCondition := domain.FilterCondition{
		Field:    "user_id",
		Operator: domain.OpEq,
		Value:    authUserId,
	}

	params.Filter.Conditions = append(params.Filter.Conditions, authCondition)

	fmt.Println(params.Filter)

	transactions, err := du.transactionRepository.Fetch(params)
	if err != nil {
		return nil, err
	}

	fmt.Println(transactions)
	// TODO: Здесь должна быть логика расчета детализации на основе полученных транзакций
	return domain.NewFinanceDetailing(time.Time{}, time.Time{}, 0, 0, 0, 0), nil
}
