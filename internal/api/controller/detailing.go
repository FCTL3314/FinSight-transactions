package controller

import (
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
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
	params, err := getParams(c, tc.cfg.Pagination.TransactionLimit)
	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(UserIDContextKey)

	responseDetailing, err := tc.usecase.Get(authUserId, &params)

	if err != nil {
		tc.errorHandler.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, responseDetailing)
}
