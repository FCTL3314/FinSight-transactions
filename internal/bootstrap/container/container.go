package container

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	Config      *config.Config
	LoggerGroup *logging.LoggerGroup
	ErrorMapper errormapper.Chain

	Transaction *TransactionContainer
}

func NewContainer(
	db *gorm.DB,
	cfg *config.Config,
	loggerGroup *logging.LoggerGroup,
) *Container {
	container := &Container{
		DB:          db,
		Config:      cfg,
		LoggerGroup: loggerGroup,
	}

	errorMapper := errormapper.BuildAllErrorsMapperChain()

	container.ErrorMapper = errorMapper
	container.Transaction = NewTransactionContainer(
		db,
		cfg,
		errorMapper,
		controller.DefaultErrorHandler(),
		loggerGroup.Transaction,
	)

	return container
}
