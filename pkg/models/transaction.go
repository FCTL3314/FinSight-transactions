package models

import (
	"time"
)

type Transaction struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Amount     float64   `json:"amount" gorm:"type:numeric(12,2);not null"`
	Name       string    `json:"name" gorm:"type:text;not null"`
	Note       string    `json:"note" gorm:"type:text"`
	CategoryID int64     `json:"category_id" gorm:"not null"`
	UserID     int64     `json:"user_id" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"not null;default:now()"`

	RecurringTransaction *RecurringTransaction `json:"recurring_transaction,omitempty" gorm:"foreignKey:TransactionID"`
}

func (t *Transaction) ToResponseTransaction() *ResponseTransaction {
	return &ResponseTransaction{
		ID:         t.ID,
		Amount:     t.Amount,
		Name:       t.Name,
		Note:       t.Note,
		CategoryID: t.CategoryID,
		UserID:     t.UserID,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func (t *Transaction) ApplyUpdateTransaction(req *UpdateTransactionRequest) {
	if req.Amount != nil {
		t.Amount = *req.Amount
	}
	if req.Name != nil {
		t.Name = *req.Name
	}
	if req.Note != nil {
		t.Note = *req.Note
	}
	if req.CategoryID != nil {
		t.CategoryID = *req.CategoryID
	}
}

func ToResponseTransactions(transactions []*Transaction) []*ResponseTransaction {
	responseTransactions := make([]*ResponseTransaction, len(transactions))
	for i, transaction := range transactions {
		responseTransactions[i] = transaction.ToResponseTransaction()
	}
	return responseTransactions
}

type ResponseTransaction struct {
	ID         uint      `json:"id"`
	Amount     float64   `json:"amount"`
	Name       string    `json:"name"`
	Note       string    `json:"note,omitempty"`
	CategoryID int64     `json:"category_id"`
	UserID     int64     `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateTransactionRequest struct {
	Amount     float64 `json:"amount" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Note       string  `json:"note"`
	CategoryID int64   `json:"category_id" binding:"required"`
	UserID     int64   `json:"user_id"`
}

func (ct *CreateTransactionRequest) ToFullTransaction() *Transaction {
	return &Transaction{
		Amount:     ct.Amount,
		Name:       ct.Name,
		Note:       ct.Note,
		CategoryID: ct.CategoryID,
		UserID:     ct.UserID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type UpdateTransactionRequest struct {
	Amount     *float64 `json:"amount"`
	Name       *string  `json:"name"`
	Note       *string  `json:"note"`
	CategoryID *int64   `json:"category_id"`
}
