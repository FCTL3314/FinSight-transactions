package controller

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
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

type transactionController struct {
	usecase      usecase.TransactionUsecase
	errorHandler *errorhandler.ErrorHandler
	Logger       logging.Logger
	cfg          *config.Config
}

func NewTransactionController(
	usecase usecase.TransactionUsecase,
	errorHandler *errorhandler.ErrorHandler,
	logger logging.Logger,
	cfg *config.Config,
) TransactionController {
	return &transactionController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (tc *transactionController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	transaction, err := tc.usecase.GetById(id)

	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := transaction.ToResponseTransaction()

	c.JSON(http.StatusOK, responseTransaction)
}

func (tc *transactionController) List(c *gin.Context) {
	params, err := getParams(c, tc.cfg.Pagination.TransactionLimit)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	paginatedResult, err := tc.usecase.List(&params)
	if err != nil {
		tc.errorHandler.Handle(c, err)
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

func (tc *transactionController) Create(c *gin.Context) {
	var transaction models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	createdTransaction, err := tc.usecase.Create(authUserId, &transaction)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := createdTransaction.ToResponseTransaction()

	c.JSON(http.StatusCreated, responseTransaction)
}

func (tc *transactionController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	var transaction models.UpdateTransactionRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	updatedTransaction, err := tc.usecase.Update(authUserId, id, &transaction)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseTransaction := updatedTransaction.ToResponseTransaction()

	c.JSON(http.StatusOK, responseTransaction)
}

func (tc *transactionController) Delete(c *gin.Context) {
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
