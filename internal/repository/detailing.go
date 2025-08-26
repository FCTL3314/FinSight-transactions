package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/Masterminds/squirrel"
)

type DetailingRepository interface {
	Get(filterParams *domain.FilterParams) (*domain.FinanceDetailing, error)
	Fetch(params *domain.Params) ([]*domain.FinanceDetailing, error)
	Count(filterParams *domain.FilterParams) (int64, error)
	Create(detailing *domain.FinanceDetailing, filterParams *domain.FilterParams) (*domain.FinanceDetailing, error)
	Update(detailing *domain.FinanceDetailing, filterParams *domain.FilterParams) (*domain.FinanceDetailing, error)
	Delete(id int64) error
}

type DefaultDetailingRepository struct {
	db *sql.DB
	sq squirrel.StatementBuilderType
}

func NewDefaultDetailingRepository(db *sql.DB) *DefaultDetailingRepository {
	return &DefaultDetailingRepository{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *DefaultDetailingRepository) scanRow(row squirrel.RowScanner) (*domain.FinanceDetailing, error) {
	var d domain.FinanceDetailing
	err := row.Scan(
		&d.ID,
		&d.UserID,
		&d.DateFrom,
		&d.DateTo,
		&d.InitialAmount,
		&d.CurrentAmount,
		&d.TotalIncome,
		&d.TotalExpense,
		&d.ProfitEstimated,
		&d.ProfitReal,
		&d.AfterAmountNet,
		&d.AfterAmountGross,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrObjectNotFound
		}
		return nil, err
	}
	return &d, nil
}

func (r *DefaultDetailingRepository) Get(filterParams *domain.FilterParams) (*domain.FinanceDetailing, error) {
	queryBuilder := r.sq.Select(
		"id",
		"user_id",
		"date_from",
		"date_to",
		"initial_amount",
		"current_amount",
		"total_income",
		"total_expense",
		"profit_estimated",
		"profit_real",
		"after_amount_net",
		"after_amount_gross",
	).From("finance_detailing").Limit(1)

	queryBuilder = applyFilters(queryBuilder, filterParams.Conditions)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	return r.scanRow(row)
}

func (r *DefaultDetailingRepository) Fetch(params *domain.Params) ([]*domain.FinanceDetailing, error) {
	queryBuilder := r.sq.Select(
		"id",
		"user_id",
		"date_from",
		"date_to",
		"initial_amount",
		"current_amount",
		"total_income",
		"total_expense",
		"profit_estimated",
		"profit_real",
		"after_amount_net",
		"after_amount_gross",
	).From("finance_detailing")

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

	detailings := make([]*domain.FinanceDetailing, 0)
	for rows.Next() {
		d, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		detailings = append(detailings, d)
	}

	return detailings, nil
}

func (r *DefaultDetailingRepository) Count(filterParams *domain.FilterParams) (int64, error) {
	queryBuilder := r.sq.Select("COUNT(*)").From("finance_detailing")

	if filterParams != nil {
		queryBuilder = applyFilters(queryBuilder, filterParams.Conditions)
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

func (r *DefaultDetailingRepository) calculateTransactionTotals(filterParams *domain.FilterParams) (float64, float64, error) {
	var totalIncome, totalExpense float64

	baseQuery := r.sq.Select("COALESCE(SUM(amount), 0)").From("transactions")
	if filterParams != nil {
		baseQuery = applyFilters(baseQuery, filterParams.Conditions)
	}

	incomeQuery, incomeArgs, err := baseQuery.Where("amount > 0").ToSql()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to build income query: %w", err)
	}

	expenseQuery, expenseArgs, err := baseQuery.Where("amount < 0").ToSql()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to build expense query: %w", err)
	}

	if err := r.db.QueryRow(incomeQuery, incomeArgs...).Scan(&totalIncome); err != nil {
		return 0, 0, fmt.Errorf("failed to query total income: %w", err)
	}

	if err := r.db.QueryRow(expenseQuery, expenseArgs...).Scan(&totalExpense); err != nil {
		return 0, 0, fmt.Errorf("failed to query total expense: %w", err)
	}

	return totalIncome, -totalExpense, nil
}

func (r *DefaultDetailingRepository) Create(detailing *domain.FinanceDetailing, filterParams *domain.FilterParams) (*domain.FinanceDetailing, error) {
	totalIncome, totalExpense, err := r.calculateTransactionTotals(filterParams)
	if err != nil {
		return nil, err
	}
	detailing.TotalIncome = totalIncome
	detailing.TotalExpense = totalExpense
	detailing.Calculate()

	queryBuilder := r.sq.Insert("finance_detailing").
		Columns(
			"user_id", "date_from", "date_to", "initial_amount", "current_amount",
			"total_income", "total_expense", "profit_estimated", "profit_real",
			"after_amount_net", "after_amount_gross",
		).
		Values(
			detailing.UserID, detailing.DateFrom, detailing.DateTo, detailing.InitialAmount, detailing.CurrentAmount,
			detailing.TotalIncome, detailing.TotalExpense, detailing.ProfitEstimated, detailing.ProfitReal,
			detailing.AfterAmountNet, detailing.AfterAmountGross,
		).
		Suffix("RETURNING id, user_id, date_from, date_to, initial_amount, current_amount, total_income, total_expense, profit_estimated, profit_real, after_amount_net, after_amount_gross")

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	return r.scanRow(row)
}

func (r *DefaultDetailingRepository) Update(detailing *domain.FinanceDetailing, filterParams *domain.FilterParams) (*domain.FinanceDetailing, error) {
	totalIncome, totalExpense, err := r.calculateTransactionTotals(filterParams)
	if err != nil {
		return nil, err
	}
	detailing.TotalIncome = totalIncome
	detailing.TotalExpense = totalExpense
	detailing.Calculate()

	sqlQuery, args, err := r.sq.Update("finance_detailing").
		Set("date_from", detailing.DateFrom).
		Set("date_to", detailing.DateTo).
		Set("initial_amount", detailing.InitialAmount).
		Set("current_amount", detailing.CurrentAmount).
		Set("total_income", detailing.TotalIncome).
		Set("total_expense", detailing.TotalExpense).
		Set("profit_estimated", detailing.ProfitEstimated).
		Set("profit_real", detailing.ProfitReal).
		Set("after_amount_net", detailing.AfterAmountNet).
		Set("after_amount_gross", detailing.AfterAmountGross).
		Where(squirrel.Eq{"id": detailing.ID}).
		Suffix("RETURNING id, user_id, date_from, date_to, initial_amount, current_amount, total_income, total_expense, profit_estimated, profit_real, after_amount_net, after_amount_gross").
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(sqlQuery, args...)
	return r.scanRow(row)
}

func (r *DefaultDetailingRepository) Delete(id int64) error {
	sqlQuery, args, err := r.sq.Delete("finance_detailing").
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
