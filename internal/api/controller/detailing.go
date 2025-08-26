package controller

import (
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/FCTL3314/FinSight-transactions/pkg/schemas"
	"github.com/gin-gonic/gin"
)

type DetailingController interface {
	GetController
	ListController
	CreateController
	UpdateController
	DeleteController
}

type detailingController struct {
	usecase      usecase.DetailingUsecase
	errorHandler *errorhandler.ErrorHandler
	Logger       logging.Logger
	cfg          *config.Config
}

func NewDetailingController(
	usecase usecase.DetailingUsecase,
	errorHandler *errorhandler.ErrorHandler,
	logger logging.Logger,
	cfg *config.Config,
) DetailingController {
	return &detailingController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (tc *detailingController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	financeDetailing, err := tc.usecase.Get(authUserId, id)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseFinanceDetailing := schemas.NewResponseFinanceDetailing(financeDetailing)
	c.JSON(http.StatusOK, responseFinanceDetailing)
}

func (tc *detailingController) List(c *gin.Context) {
	params, err := getParams(
		c,
		tc.cfg.Pagination.FinanceDetailingLimit,
		WithFilters(
			map[string][]string{
				"date_from": {"gte", "lte", "gt", "lt"},
				"date_to":   {"gte", "lte", "gt", "lt"},
			},
		),
		WithOrdering("created_at DESC"),
	)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)
	paginatedResult, err := tc.usecase.List(authUserId, &params)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseDetailing := schemas.NewResponseFinanceDetailingList(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse[*schemas.ResponseFinanceDetailing]{
		Count:   paginatedResult.Count,
		Limit:   params.Pagination.Limit,
		Offset:  params.Pagination.Offset,
		Results: responseDetailing,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (tc *detailingController) Create(c *gin.Context) {
	var req schemas.CreateFinanceDetailingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	createdDetailing, err := tc.usecase.Create(authUserId, &req)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseFinanceDetailing := schemas.NewResponseFinanceDetailing(createdDetailing)
	c.JSON(http.StatusCreated, responseFinanceDetailing)
}

func (tc *detailingController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	var req schemas.UpdateFinanceDetailingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	updatedDetailing, err := tc.usecase.Update(authUserId, id, &req)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseFinanceDetailing := schemas.NewResponseFinanceDetailing(updatedDetailing)
	c.JSON(http.StatusOK, responseFinanceDetailing)
}

func (tc *detailingController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	if err := tc.usecase.Delete(authUserId, id); err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
