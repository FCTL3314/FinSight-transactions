package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase/access"
	"github.com/FCTL3314/FinSight-transactions/pkg/schemas"
)

type TransactionUsecase interface {
	GetById(id int64) (*domain.Transaction, error)
	Get(params *domain.FilterParams) (*domain.Transaction, error)
	List(params *domain.Params) (*domain.PaginatedResult[*domain.Transaction], error)
	Create(authUserId int64, createTransactionRequest *schemas.CreateTransactionRequest) (*domain.Transaction, error)
	Update(authUserId, id int64, updateTransactionRequest *schemas.UpdateTransactionRequest) (*domain.Transaction, error)
	Delete(authUserId, id int64) error
}

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
	accessPolicy          access.Policy[domain.Transaction]
	cfg                   *config.Config
}

func NewTransactionUsecase(
	transactionRepository repository.TransactionRepository,
	accessPolicy access.Policy[domain.Transaction],
	cfg *config.Config,
) TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionRepository,
		accessPolicy:          accessPolicy,
		cfg:                   cfg,
	}
}

func (wu *transactionUsecase) GetById(id int64) (*domain.Transaction, error) {
	transaction, err := wu.transactionRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (wu *transactionUsecase) Get(params *domain.FilterParams) (*domain.Transaction, error) {
	transaction, err := wu.transactionRepository.Get(params)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (wu *transactionUsecase) List(params *domain.Params) (*domain.PaginatedResult[*domain.Transaction], error) {
	transactions, err := wu.transactionRepository.Fetch(params)
	if err != nil {
		return nil, err
	}

	count, err := wu.transactionRepository.Count(params.Filter)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResult[*domain.Transaction]{Results: transactions, Count: count}, nil
}

func (wu *transactionUsecase) Create(authUserId int64, createTransactionRequest *schemas.CreateTransactionRequest) (*domain.Transaction, error) {
	transaction := createTransactionRequest.ToDomainModel(authUserId)
	return wu.transactionRepository.Create(transaction)
}

func (wu *transactionUsecase) Update(authUserId int64, id int64, updateTransactionRequest *schemas.UpdateTransactionRequest) (*domain.Transaction, error) {
	transactionToUpdate, err := wu.transactionRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if !wu.accessPolicy.HasAccess(authUserId, transactionToUpdate) {
		return nil, domain.ErrAccessDenied
	}

	updateTransactionRequest.ApplyToDomainModel(transactionToUpdate)

	updatedTransaction, err := wu.transactionRepository.Update(transactionToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedTransaction, nil
}

func (wu *transactionUsecase) Delete(authUserId int64, id int64) error {
	transaction, err := wu.transactionRepository.GetById(id)
	if err != nil {
		return err
	}

	if !wu.accessPolicy.HasAccess(authUserId, transaction) {
		return domain.ErrAccessDenied
	}

	return wu.transactionRepository.Delete(id)
}
