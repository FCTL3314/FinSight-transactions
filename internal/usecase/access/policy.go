package access

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
)

type Policy[T any] interface {
	HasAccess(authUserID int64, resource *T) bool
}

type transactionAccessPolicy struct{}

func NewTransactionAccessPolicy() Policy[domain.Transaction] {
	return &transactionAccessPolicy{}
}

func (wa *transactionAccessPolicy) HasAccess(authUserID int64, transaction *domain.Transaction) bool {
	return authUserID == transaction.UserID
}
