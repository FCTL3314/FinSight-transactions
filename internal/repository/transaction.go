package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/Masterminds/squirrel"
)

type TransactionRepository interface {
	Repository[domain.Transaction]
}

type DefaultTransactionRepository struct {
	db *sql.DB
	sq squirrel.StatementBuilderType
}

func NewDefaultTransactionRepository(db *sql.DB) *DefaultTransactionRepository {
	return &DefaultTransactionRepository{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *DefaultTransactionRepository) scanRow(row squirrel.RowScanner) (*domain.Transaction, error) {
	var t domain.Transaction
	err := row.Scan(
		&t.ID,
		&t.Amount,
		&t.Name,
		&t.Note,
		&t.CategoryID,
		&t.UserID,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}
	return &t, nil
}

func (r *DefaultTransactionRepository) GetById(id int64) (*domain.Transaction, error) {
	return r.Get(&domain.FilterParams{
		Conditions: []domain.FilterCondition{
			{Field: "id", Operator: domain.OpEq, Value: id},
		},
	})
}

func (r *DefaultTransactionRepository) Get(filterParams *domain.FilterParams) (*domain.Transaction, error) {
	queryBuilder := r.sq.Select(
		"id",
		"amount",
		"name",
		"note",
		"category_id",
		"user_id",
		"created_at",
		"updated_at",
	).From("transactions").Limit(1)

	queryBuilder = applyFilters(queryBuilder, filterParams.Conditions)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	return r.scanRow(row)
}

func (r *DefaultTransactionRepository) Fetch(params *domain.Params) ([]*domain.Transaction, error) {
	queryBuilder := r.sq.Select(
		"id",
		"amount",
		"name",
		"note",
		"category_id",
		"user_id",
		"created_at",
		"updated_at",
	).From("transactions")

	if params.Filter != nil {
		queryBuilder = applyFilters(queryBuilder, params.Filter.Conditions)
	}
	if params.OrderParams != nil && params.OrderParams.Order != "" {
		queryBuilder = queryBuilder.OrderBy(params.OrderParams.Order)
	}
	if params.Pagination != nil {
		queryBuilder = queryBuilder.Limit(uint64(params.Pagination.Limit)).Offset(uint64(params.Pagination.Offset))
	}

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
			// TODO: Pass logger and log error
		}
	}(rows)

	transactions := make([]*domain.Transaction, 0)
	for rows.Next() {
		t, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}

func (r *DefaultTransactionRepository) Create(transaction *domain.Transaction) (*domain.Transaction, error) {
	sqlQuery, args, err := r.sq.Insert("transactions").
		Columns(
			"amount",
			"name",
			"note",
			"category_id",
			"user_id",
			"created_at",
			"updated_at",
		).
		Values(transaction.Amount, transaction.Name, transaction.Note, transaction.CategoryID, transaction.UserID, transaction.CreatedAt, transaction.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(sqlQuery, args...).Scan(&transaction.ID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *DefaultTransactionRepository) Update(transaction *domain.Transaction) (*domain.Transaction, error) {
	sqlQuery, args, err := r.sq.Update("transactions").
		Set("amount", transaction.Amount).
		Set("name", transaction.Name).
		Set("note", transaction.Note).
		Set("category_id", transaction.CategoryID).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": transaction.ID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	result, err := r.db.Exec(sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, domain.ErrObjectNotFound
	}

	return transaction, nil
}

func (r *DefaultTransactionRepository) Delete(id int64) error {
	sqlQuery, args, err := r.sq.Delete("transactions").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	result, err := r.db.Exec(sqlQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrObjectNotFound
	}

	return nil
}

func (r *DefaultTransactionRepository) Count(params *domain.FilterParams) (int64, error) {
	queryBuilder := r.sq.Select("COUNT(*)").From("transactions")

	if params != nil {
		queryBuilder = applyFilters(queryBuilder, params.Conditions)
	}

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var count int64
	err = r.db.QueryRow(sqlQuery, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
