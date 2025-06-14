package dependency

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/database"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContainer struct {
	GinEngine    *gin.Engine
	BaseRouter   *gin.RouterGroup
	DB           *gorm.DB
	Config       *config.Config
	LoggersGroup *logging.LoggersGroup

	*TransactionContainer
}

func NewAppContainer() *AppContainer {
	var c AppContainer

	c.setupGin()
	c.setupConfig()
	c.setupLoggers()
	c.setupDatabase()
	c.setupTransaction()

	return &c
}

func (c *AppContainer) setupGin() {
	engine := gin.Default()
	router := engine.Group("/api/v1/")

	c.GinEngine = engine
	c.BaseRouter = router
}

func (c *AppContainer) setupConfig() {
	cfg, err := config.Load()
	if err != nil {
		c.LoggersGroup.General.Fatal(
			"Failed to load configuration. Please check environmental files",
			logging.WithError(err),
		)
	}
	c.Config = cfg
}

func (c *AppContainer) setupLoggers() {
	generalLogger := logging.InitGeneralLogger()
	transactionLogger := logging.InitTransactionLogger()

	c.LoggersGroup = logging.NewLoggersGroup(
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
		c.LoggersGroup.General.Fatal("database connection failed", logging.WithError(err))
	}
	c.DB = db
}

func (c *AppContainer) setupTransaction() {
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	errorHandler := errorhandler.NewErrorHandler(errorMapper)
	errorhandler.RegisterAllErrorHandlers(errorHandler)

	c.TransactionContainer = NewTransactionContainer(
		c.BaseRouter,
		c.DB,
		c.Config,
		errorHandler,
		c.LoggersGroup.Transaction,
	)
}
