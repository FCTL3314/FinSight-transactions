package repository

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Repository[models.Transaction]
}

type DefaultTransactionRepository struct {
	db        *gorm.DB
	toPreload []string
}

func NewDefaultTransactionRepository(db *gorm.DB) *DefaultTransactionRepository {
	return &DefaultTransactionRepository{db: db}
}

func (wr *DefaultTransactionRepository) GetById(id int64) (*models.Transaction, error) {
	return wr.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (wr *DefaultTransactionRepository) Get(filterParams *domain.FilterParams) (*models.Transaction, error) {
	var transaction models.Transaction
	query := wr.db.Where(filterParams.Query, filterParams.Args...)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if err := (query.First(&transaction)).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (wr *DefaultTransactionRepository) Fetch(params *domain.Params) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&transactions)).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (wr *DefaultTransactionRepository) Create(transaction *models.Transaction) (*models.Transaction, error) {
	if err := (wr.db.Save(&transaction)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.Transaction{}), wr.toPreload)
	if err := query.First(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (wr *DefaultTransactionRepository) Update(transaction *models.Transaction) (*models.Transaction, error) {
	if err := (wr.db.Save(&transaction)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.Transaction{}), wr.toPreload)
	if err := query.First(transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (wr *DefaultTransactionRepository) Delete(id int64) error {
	result := wr.db.Where("id = ?", id).Delete(&models.Transaction{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (wr *DefaultTransactionRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (wr.db.Model(&models.Transaction{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
