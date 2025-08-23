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

func getFilterParams(c *gin.Context, allowedFilterFields []string) (*domain.FilterParams, error) {
	queryParams := c.Request.URL.Query()
	filter := domain.NewFilterParams()

	for key, values := range queryParams {
		if len(values) == 0 || values[0] == "" {
			continue
		}

		fieldName := key
		operator := domain.OpEq

		if parts := strings.Split(key, "_"); len(parts) > 1 {
			suffix := parts[len(parts)-1]
			if op, ok := operatorMap[suffix]; ok {
				fieldName = strings.Join(parts[:len(parts)-1], "_")
				operator = op
			}
		}

		if !slice.Contains(allowedFilterFields, fieldName) {
			continue
		}

		var val interface{} = values[0]
		if operator == domain.OpIn {
			val = strings.Split(values[0], ",")
		}

		condition := domain.FilterCondition{
			Field:    fieldName,
			Operator: operator,
			Value:    val,
		}
		filter.Conditions = append(filter.Conditions, condition)
	}

	return filter, nil
}

func getParams(c *gin.Context, paginationMaxLimit int, allowedFilterFields ...string) (domain.Params, error) {
	paginationParams, err := getPaginationParams(c, paginationMaxLimit)
	if err != nil {
		return domain.Params{}, err
	}

	filterParams, err := getFilterParams(c, allowedFilterFields)
	if err != nil {
		return domain.Params{}, err
	}

	orderParams := domain.OrderParams{Order: "created_at DESC"}

	return domain.Params{
		Pagination:  &paginationParams,
		Filter:      filterParams,
		OrderParams: orderParams,
	}, nil
}
