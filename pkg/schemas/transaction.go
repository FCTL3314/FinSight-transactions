package schemas

import (
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
)

type CreateTransactionRequest struct {
	Amount     float64   `json:"amount" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Note       string    `json:"note"`
	CategoryID int64     `json:"category_id" binding:"required"`
	MadeAt     time.Time `json:"made_at"`
}

func (req *CreateTransactionRequest) ToDomainModel(userID int64) *domain.Transaction {
	madeAt := req.MadeAt
	if madeAt.IsZero() {
		madeAt = time.Now()
	}

	return &domain.Transaction{
		Amount:     req.Amount,
		Name:       req.Name,
		Note:       req.Note,
		CategoryID: req.CategoryID,
		UserID:     userID,
		MadeAt:     madeAt,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type UpdateTransactionRequest struct {
	Amount     *float64   `json:"amount"`
	Name       *string    `json:"name"`
	Note       *string    `json:"note"`
	CategoryID *int64     `json:"category_id"`
	MadeAt     *time.Time `json:"made_at"`
}

func (req *UpdateTransactionRequest) ApplyToDomainModel(t *domain.Transaction) {
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
	if req.MadeAt != nil {
		t.MadeAt = *req.MadeAt
	}
}

type ResponseTransaction struct {
	ID         uint      `json:"id"`
	Amount     float64   `json:"amount"`
	Name       string    `json:"name"`
	Note       string    `json:"note,omitempty"`
	CategoryID int64     `json:"category_id"`
	UserID     int64     `json:"user_id"`
	MadeAt     time.Time `json:"made_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewResponseTransaction(t *domain.Transaction) *ResponseTransaction {
	if t == nil {
		return nil
	}

	return &ResponseTransaction{
		ID:         t.ID,
		Amount:     t.Amount,
		Name:       t.Name,
		Note:       t.Note,
		CategoryID: t.CategoryID,
		UserID:     t.UserID,
		MadeAt:     t.MadeAt,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
	}
}

func NewResponseTransactionList(transactions []*domain.Transaction) []*ResponseTransaction {
	response := make([]*ResponseTransaction, len(transactions))
	for i, transaction := range transactions {
		response[i] = NewResponseTransaction(transaction)
	}
	return response
}
