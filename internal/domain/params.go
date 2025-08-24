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

func NewFilterParams(conditions ...FilterCondition) *FilterParams {
	return &FilterParams{Conditions: conditions}
}

type PaginationParams struct {
	Limit  int
	Offset int
}

type OrderParams struct {
	Order string
}

func NewOrderParams(order string) *OrderParams {
	return &OrderParams{Order: order}
}

type Params struct {
	Filter     *FilterParams
	Pagination *PaginationParams
	*OrderParams
}

func NewParams(filter *FilterParams, pagination *PaginationParams, order *OrderParams) *Params {
	return &Params{
		Filter:      filter,
		Pagination:  pagination,
		OrderParams: order,
	}
}
