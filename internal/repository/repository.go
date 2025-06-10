package repository

import "github.com/FCTL3314/FinSight-transactions/internal/domain"

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
