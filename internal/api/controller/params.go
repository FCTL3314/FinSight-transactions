package controller

import (
	"strconv"
	"strings"

	"github.com/FCTL3314/FinSight-transactions/internal/collections/slice"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/gin-gonic/gin"
)

const (
	UserIDContextKey = "userID"
)

var (
	operatorMap = map[string]string{
		"gt":   domain.OpGt,
		"gte":  domain.OpGte,
		"lt":   domain.OpLt,
		"lte":  domain.OpLte,
		"neq":  domain.OpNotEq,
		"like": domain.OpLike,
		"in":   domain.OpIn,
	}
)

func getParamAsInt64(c *gin.Context, key string) (int64, error) {
	id, err := strconv.ParseInt(c.Param(key), 10, 64)
	if err != nil {
		return 0, &domain.ErrInvalidURLParam{Param: key}
	}
	return id, nil
}

func getPaginationParams(c *gin.Context, maxLimit int) (domain.PaginationParams, error) {
	limitStr := c.DefaultQuery("limit", strconv.Itoa(maxLimit))
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	} else if limit > maxLimit {
		return domain.PaginationParams{}, &domain.ErrPaginationLimitExceeded{MaxLimit: maxLimit}
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	return domain.PaginationParams{
		Limit:  limit,
		Offset: offset,
	}, nil
}

func getFilterParams(c *gin.Context, allowedFilters map[string][]string) (*domain.FilterParams, error) {
	queryParams := c.Request.URL.Query()
	filter := domain.NewFilterParams()

	if allowedFilters == nil {
		return filter, nil
	}

	for key, values := range queryParams {
		if len(values) == 0 || values[0] == "" {
			continue
		}

		fieldName := key
		operatorSuffix := "eq"

		if parts := strings.Split(key, "_"); len(parts) > 1 {
			lastPart := parts[len(parts)-1]
			if _, ok := operatorMap[lastPart]; ok {
				fieldName = strings.Join(parts[:len(parts)-1], "_")
				operatorSuffix = lastPart
			}
		}

		allowedOperators, fieldIsAllowed := allowedFilters[fieldName]
		if !fieldIsAllowed {
			continue
		}

		if !slice.Contains(allowedOperators, operatorSuffix) {
			continue
		}

		finalOperator := operatorMap[operatorSuffix]
		if operatorSuffix == "eq" {
			finalOperator = domain.OpEq
		}

		var val interface{} = values[0]
		if finalOperator == domain.OpIn {
			val = strings.Split(values[0], ",")
		}

		condition := domain.FilterCondition{
			Field:    fieldName,
			Operator: finalOperator,
			Value:    val,
		}
		filter.Conditions = append(filter.Conditions, condition)
	}

	return filter, nil
}

type paramsConfig struct {
	paginationMaxLimit  int
	allowedFilterFields map[string][]string
	defaultOrder        string
}

type ParamOption func(*paramsConfig)

func WithAllowedFilters(filters map[string][]string) ParamOption {
	return func(cfg *paramsConfig) {
		cfg.allowedFilterFields = filters
	}
}

func WithDefaultOrder(order string) ParamOption {
	return func(cfg *paramsConfig) {
		cfg.defaultOrder = order
	}
}

func getParams(c *gin.Context, paginationMaxLimit int, opts ...ParamOption) (domain.Params, error) {
	cfg := &paramsConfig{
		paginationMaxLimit: paginationMaxLimit,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	paginationParams, err := getPaginationParams(c, cfg.paginationMaxLimit)
	if err != nil {
		return domain.Params{}, err
	}

	filterParams, err := getFilterParams(c, cfg.allowedFilterFields)
	if err != nil {
		return domain.Params{}, err
	}

	orderParams := domain.OrderParams{Order: cfg.defaultOrder}

	return domain.Params{
		Pagination:  &paginationParams,
		Filter:      filterParams,
		OrderParams: orderParams,
	}, nil
}
