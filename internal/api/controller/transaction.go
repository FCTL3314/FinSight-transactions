package controller

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController interface {
	Controller
}

type DefaultTransactionController struct {
	usecase      usecase.TransactionUsecase
	errorHandler *ErrorHandler
	Logger       logging.Logger
	cfg          *config.Config
}

func NewDefaultTransactionController(
	usecase usecase.TransactionUsecase,
	errorHandler *ErrorHandler,
	logger logging.Logger,
	cfg *config.Config,
) *DefaultTransactionController {
	return &DefaultTransactionController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (wc *DefaultTransactionController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	transaction, err := wc.usecase.GetById(id)

	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := transaction.ToResponseTransaction()

	c.JSON(http.StatusOK, responseTransaction)
}

func (wc *DefaultTransactionController) List(c *gin.Context) {
	params, err := getParams(c, wc.cfg.Pagination.TransactionLimit)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	paginatedResult, err := wc.usecase.List(&params)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseTransactions := models.ToResponseTransactions(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse[*models.ResponseTransaction]{
		Count:   paginatedResult.Count,
		Limit:   params.Pagination.Limit,
		Offset:  params.Pagination.Offset,
		Results: responseTransactions,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (wc *DefaultTransactionController) Create(c *gin.Context) {
	var transaction models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	createdTransaction, err := wc.usecase.Create(authUserId, &transaction)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := createdTransaction.ToResponseTransaction()

	c.JSON(http.StatusCreated, responseTransaction)
}

func (wc *DefaultTransactionController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var transaction models.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	updatedTransaction, err := wc.usecase.Update(authUserId, id, &transaction)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := updatedTransaction.ToResponseTransaction()

	c.JSON(http.StatusOK, responseTransaction)
}

func (wc *DefaultTransactionController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	if err := wc.usecase.Delete(authUserId, id); err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
