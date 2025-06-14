package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
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
	cfg                   *config.Config
}

func NewDefaultTransactionUsecase(
	transactionRepository repository.TransactionRepository,
	cfg *config.Config,
) *DefaultTransactionUsecase {
	return &DefaultTransactionUsecase{
		transactionRepository: transactionRepository,
		cfg:                   cfg,
	}
}

func (wu *DefaultTransactionUsecase) GetById(id int64) (*models.Transaction, error) {
	transaction, err := wu.transactionRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (wu *DefaultTransactionUsecase) Get(params *domain.FilterParams) (*models.Transaction, error) {
	transaction, err := wu.transactionRepository.Get(params)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (wu *DefaultTransactionUsecase) List(params *domain.Params) (*domain.PaginatedResult[*models.Transaction], error) {
	transactions, err := wu.transactionRepository.Fetch(params)
	if err != nil {
		return nil, err
	}

	count, err := wu.transactionRepository.Count(&domain.FilterParams{})
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResult[*models.Transaction]{Results: transactions, Count: count}, nil
}

func (wu *DefaultTransactionUsecase) Create(authUserId int64, createTransactionRequest *models.CreateTransactionRequest) (*models.Transaction, error) {
	transaction := createTransactionRequest.ToFullTransaction()
	return wu.transactionRepository.Create(transaction)
}

func (wu *DefaultTransactionUsecase) Update(authUserId int64, id int64, updateTransactionRequest *models.UpdateTransactionRequest) (*models.Transaction, error) {
	transactionToUpdate, err := wu.transactionRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	//if !wu.accessManager.HasAccess(authUserId, transactionToUpdate) {
	//	return nil, domain.ErrAccessDenied
	//}

	transactionToUpdate.ApplyUpdateTransaction(updateTransactionRequest)

	updatedTransaction, err := wu.transactionRepository.Update(transactionToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedTransaction, nil
}

func (wu *DefaultTransactionUsecase) Delete(authUserId int64, id int64) error {
	//_, err := wu.transactionRepository.GetById(id)
	//if err != nil {
	//	return wu.errorMapper.MapError(err)
	//}

	//if !wu.accessManager.HasAccess(authUserId, workout) {
	//	return domain.ErrAccessDenied
	//}

	return wu.transactionRepository.Delete(id)
}
