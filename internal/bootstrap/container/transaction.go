package container

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionContainer struct {
	Repository        repository.TransactionRepository
	Usecase           usecase.TransactionUsecase
	Controller        controller.TransactionController
	Router            router.TransactionRouter
	RouterRegistrator router.Registrator
	Logger            logging.Logger
}

func NewTransactionContainer(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *config.Config,
	errorMapper errormapper.Chain,
	errorHandler *controller.ErrorHandler,
	logger logging.Logger,
) *TransactionContainer {
	var container TransactionContainer

	container.Repository = repository.NewDefaultTransactionRepository(db)
	container.Usecase = usecase.NewDefaultTransactionUsecase(
		container.Repository,
		errorMapper,
		cfg,
	)
	container.Controller = controller.NewDefaultTransactionController(
		container.Usecase,
		errorHandler,
		logger,
		cfg,
	)
	container.Router = router.NewDefaultTransactionRouter(
		baseRouter,
		container.Controller,
		cfg,
		container.Logger,
	)
	container.RouterRegistrator = router.NewTransactionRouterRegistrator(
		container.Router,
	)

	return &container
}
