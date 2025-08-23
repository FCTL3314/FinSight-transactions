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
	"github.com/gin-gonic/gin"
)

type DetailingContainer struct {
	Repository        repository.TransactionRepository
	Usecase           usecase.DetailingUsecase
	Controller        controller.DetailingController
	Router            router.DetailingRouter
	RouterRegistrator router.Registrator
	Logger            logging.Logger
}

func NewDetailingContainer(
	baseRouter *gin.RouterGroup,
	db *sql.DB,
	cfg *config.Config,
	errorHandler *errorhandler.ErrorHandler,
	logger logging.Logger,
) *DetailingContainer {
	var container DetailingContainer

	container.Repository = repository.NewDefaultTransactionRepository(db)

	container.Usecase = usecase.NewDetailingUsecase(
		container.Repository,
		cfg,
	)
	container.Controller = controller.NewDetailingController(
		container.Usecase,
		errorHandler,
		logger,
		cfg,
	)
	container.Router = router.NewDetailingRouter(
		baseRouter,
		container.Controller,
		cfg,
	)
	container.RouterRegistrator = router.NewDetailingRouterRegistrator(
		container.Router,
	)

	return &container
}
