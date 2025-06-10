package container

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/database"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContainer struct {
	GinEngine   *gin.Engine
	Router      *gin.RouterGroup
	DB          *gorm.DB
	Config      *config.Config
	LoggerGroup *logging.LoggerGroup

	Transaction *TransactionContainer
}

func NewAppContainer() *AppContainer {
	c := &AppContainer{}

	c.setupGin()
	c.setupConfig()
	c.setupLoggers()
	c.setupDatabase()
	c.setupTransaction()

	return c
}

func (c *AppContainer) setupGin() {
	engine := gin.Default()
	router := engine.Group("/api/v1/")

	c.GinEngine = engine
	c.Router = router
}

func (c *AppContainer) setupConfig() {
	cfg, err := config.Load()
	if err != nil {
		c.LoggerGroup.General.Fatal(
			"failed to load configuration. Please check environmental files",
			logging.WithError(err),
		)
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
		c.LoggerGroup.General.Fatal("database connection failed", logging.WithError(err))
	}
	c.DB = db
}

func (c *AppContainer) setupTransaction() {
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	errorHandler := controller.DefaultErrorHandler()

	c.Transaction = NewTransactionContainer(
		c.Router,
		c.DB,
		c.Config,
		errorMapper,
		errorHandler,
		c.LoggerGroup.Transaction,
	)
}
