package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
)

type TransactionUsecase interface {
	GetById(id int64) (*models.Transaction, error)
	Get(params *domain.FilterParams) (*models.Transaction, error)
	List(params *domain.Params) (*domain.PaginatedResult[*models.Transaction], error)
	Create(authUserId int64, createTransactionRequest *models.CreateTransactionRequest) (*models.Transaction, error)
	Update(authUserId, id int64, updateTransactionRequest *models.UpdateTransactionRequest) (*models.Transaction, error)
	Delete(authUserId, id int64) error
}

type DefaultTransactionUsecase struct {
	transactionRepository repository.TransactionRepository
	errorMapper           errormapper.Chain
	cfg                   *config.Config
}

func NewDefaultTransactionUsecase(
	transactionRepository repository.TransactionRepository,
	errorMapper errormapper.Chain,
	cfg *config.Config,
) *DefaultTransactionUsecase {
	return &DefaultTransactionUsecase{
		transactionRepository: transactionRepository,
		errorMapper:           errorMapper,
		cfg:                   cfg,
	}
}

func (wu *DefaultTransactionUsecase) GetById(id int64) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

func (wu *DefaultTransactionUsecase) Get(params *domain.FilterParams) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

func (wu *DefaultTransactionUsecase) List(params *domain.Params) (*domain.PaginatedResult[*models.Transaction], error) {
	return &domain.PaginatedResult[*models.Transaction]{}, nil
}

func (wu *DefaultTransactionUsecase) Create(authUserId int64, createTransactionRequest *models.CreateTransactionRequest) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

func (wu *DefaultTransactionUsecase) Update(authUserId int64, id int64, updateTransactionRequest *models.UpdateTransactionRequest) (*models.Transaction, error) {
	return &models.Transaction{}, nil
}

func (wu *DefaultTransactionUsecase) Delete(authUserId int64, id int64) error {
	return nil
}
