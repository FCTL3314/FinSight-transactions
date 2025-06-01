package repository

import (
	"context"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	Create(ctx context.Context, tx *models.CreateTransaction) error
	FindAll(ctx context.Context) ([]models.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) Create(ctx context.Context, tx *models.CreateTransaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *transactionRepo) FindAll(ctx context.Context) ([]models.Transaction, error) {
	var list []models.Transaction
	err := r.db.WithContext(ctx).Preload("Recurring").Find(&list).Error
	return list, err
}
