package container

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"gorm.io/gorm"
)

type TransactionContainer struct {
	Repository repository.TransactionRepository
	Usecase    usecase.TransactionUsecase
	Controller controller.TransactionController
}

func NewTransactionContainer(
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
	container.Controller = controller.NewTransactionController(
		container.Usecase,
		errorHandler,
		logger,
		cfg,
	)

	return &container
}
