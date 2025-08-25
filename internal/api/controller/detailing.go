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
	var req schemas.GetFinanceDetailingRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	financeDetailing, err := tc.usecase.Get(authUserId, &req)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	responseFinanceDetailing := schemas.NewResponseFinanceDetailing(financeDetailing)
	c.JSON(http.StatusOK, responseFinanceDetailing)
}
