package dependency

import (
	"database/sql"

	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase/access"
	"github.com/gin-gonic/gin"
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
	db *sql.DB,
	cfg *config.Config,
	errorHandler *errorhandler.ErrorHandler,
	logger logging.Logger,
) *TransactionContainer {
	var container TransactionContainer

	container.Repository = repository.NewDefaultTransactionRepository(db)

	accessPolicy := access.NewTransactionAccessPolicy()
	container.Usecase = usecase.NewTransactionUsecase(
		container.Repository,
		accessPolicy,
		cfg,
	)
	container.Controller = controller.NewTransactionController(
		container.Usecase,
		errorHandler,
		logger,
		cfg,
	)
	container.Router = router.NewTransactionRouter(
		baseRouter,
		container.Controller,
		cfg,
	)
	container.RouterRegistrator = router.NewTransactionRouterRegistrator(
		container.Router,
	)

	return &container
}
