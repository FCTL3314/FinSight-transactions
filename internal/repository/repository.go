package repository

import (
	"fmt"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/Masterminds/squirrel"
)

type GetterById[T any] interface {
	GetById(id int64) (*T, error)
}

type Getter[T any] interface {
	Get(filterParams *domain.FilterParams) (*T, error)
}

type Fetcher[T any] interface {
	Fetch(params *domain.Params) ([]*T, error)
}

type Creator[T any] interface {
	Create(entity *T) (*T, error)
}

type Updater[T any] interface {
	Update(entity *T) (*T, error)
}

type Deleter[T any] interface {
	Delete(id int64) error
}

type Counter[T any] interface {
	Count(params *domain.FilterParams) (int64, error)
}

type Repository[T any] interface {
	GetterById[T]
	Getter[T]
	Fetcher[T]
	Creator[T]
	Updater[T]
	Deleter[T]
	Counter[T]
}

func applyFilters(builder squirrel.SelectBuilder, conditions []domain.FilterCondition) squirrel.SelectBuilder {
	if conditions == nil {
		return builder
	}

	for _, cond := range conditions {
		switch cond.Operator {
		case domain.OpEq:
			builder = builder.Where(squirrel.Eq{cond.Field: cond.Value})
		case domain.OpNotEq:
			builder = builder.Where(squirrel.NotEq{cond.Field: cond.Value})
		case domain.OpGt:
			builder = builder.Where(squirrel.Gt{cond.Field: cond.Value})
		case domain.OpGte:
			builder = builder.Where(squirrel.GtOrEq{cond.Field: cond.Value})
		case domain.OpLt:
			builder = builder.Where(squirrel.Lt{cond.Field: cond.Value})
		case domain.OpLte:
			builder = builder.Where(squirrel.LtOrEq{cond.Field: cond.Value})
		case domain.OpLike:
			builder = builder.Where(squirrel.Like{cond.Field: fmt.Sprintf("%%%v%%", cond.Value)})
		case domain.OpIn:
			builder = builder.Where(squirrel.Eq{cond.Field: cond.Value})
		}
	}

	return builder
}
