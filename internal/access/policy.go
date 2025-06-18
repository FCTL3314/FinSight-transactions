package access

import (
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
)

type Policy[T any] interface {
	HasAccess(authUserID int64, resource T) bool
}

type transactionAccessPolicy struct{}

func NewTransactionAccessPolicy() Policy[models.Transaction] {
	return &transactionAccessPolicy{}
}

func (wa *transactionAccessPolicy) HasAccess(authUserID int64, transaction models.Transaction) bool {
	return authUserID == transaction.UserID
}
