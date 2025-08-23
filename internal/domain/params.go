package domain

const (
	OpEq    = "eq"
	OpNotEq = "neq"
	OpGt    = "gt"
	OpGte   = "gte"
	OpLt    = "lt"
	OpLte   = "lte"
	OpLike  = "like"
	OpIn    = "in"
)

type FilterCondition struct {
	Field    string
	Operator string
	Value    interface{}
}

type FilterParams struct {
	Conditions []FilterCondition
}

func NewFilterParams() *FilterParams {
	return &FilterParams{Conditions: make([]FilterCondition, 0)}
}

type PaginationParams struct {
	Limit  int
	Offset int
}

type OrderParams struct {
	Order string
}

type Params struct {
	Filter     *FilterParams
	Pagination *PaginationParams
	OrderParams
}
