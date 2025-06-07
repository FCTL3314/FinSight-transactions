package container

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/database"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type AppContainer struct {
	Router      *gin.Engine
	DB          *gorm.DB
	Config      *config.Config
	LoggerGroup *logging.LoggerGroup
	ErrorMapper errormapper.Chain
	Transaction *TransactionContainer
}

func NewAppContainer() *AppContainer {
	container := &AppContainer{}

	errorMapper := errormapper.BuildAllErrorsMapperChain()
	errorHandler := controller.DefaultErrorHandler()

	container.setupConfig()
	container.setupLoggers()
	container.setupDatabase()

	container.Router = gin.Default()
	container.ErrorMapper = errorMapper
	container.Transaction = NewTransactionContainer(
		container.DB,
		container.Config,
		errorMapper,
		errorHandler,
		container.LoggerGroup.Transaction,
	)

	return container
}

func (c *AppContainer) setupConfig() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration. Please check environmental files: ", err)
	}
	c.Config = cfg
}

func (c *AppContainer) setupLoggers() {
	generalLogger := logging.InitGeneralLogger()
	transactionLogger := logging.InitTransactionLogger()

	c.LoggerGroup = logging.NewLoggerGroup(
		generalLogger,
		transactionLogger,
	)
}

func (c *AppContainer) setupDatabase() {
	dbConnector := database.NewGormConnector(
		c.Config.Database.Name,
		c.Config.Database.User,
		c.Config.Database.Password,
		c.Config.Database.Host,
		c.Config.Database.Port,
	)

	db, err := dbConnector.Connect()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	c.DB = db
}
