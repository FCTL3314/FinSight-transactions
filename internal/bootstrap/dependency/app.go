package dependency

import (
	"database/sql"

	"github.com/FCTL3314/FinSight-transactions/internal/api/controller/errorhandler"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/database"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type AppContainer struct {
	GinEngine    *gin.Engine
	BaseRouter   *gin.RouterGroup
	V1Router     *gin.RouterGroup
	DB           *sql.DB
	Config       *config.Config
	LoggersGroup *logging.LoggersGroup

	SystemContainer      *SystemContainer
	TransactionContainer *TransactionContainer
	DetailingContainer   *DetailingContainer
}

func NewAppContainer() *AppContainer {
	var c AppContainer

	c.setupGin()
	c.setupLoggers()
	c.setupConfig()
	c.setupDatabase()
	c.setupHealthCheck()
	c.setupTransaction()
	c.setupDetailing()

	return &c
}

func (c *AppContainer) setupGin() {
	engine := gin.Default()
	baseRouter := engine.Group("/")
	v1Router := baseRouter.Group("/api/v1/")

	c.GinEngine = engine
	c.BaseRouter = baseRouter
	c.V1Router = v1Router
}

func (c *AppContainer) setupConfig() {
	cfg, err := config.Load()
	if err != nil {
		c.LoggersGroup.General.Fatal(
			"Failed to load configuration. Please check environmental files. "+
				"If you run the application via Docker, "+
				"check if you use the WORKDIR directive.",
			logging.WithError(err),
		)
	}
	c.Config = cfg
}

func (c *AppContainer) setupLoggers() {
	generalLogger := logging.InitGeneralLogger()
	transactionLogger := logging.InitTransactionLogger()
	detailingLogger := logging.InitDetailingLogger()

	c.LoggersGroup = logging.NewLoggersGroup(
		generalLogger,
		transactionLogger,
		detailingLogger,
	)
}

func (c *AppContainer) setupDatabase() {
	dbConnector := database.NewPgxConnector(
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

func (c *AppContainer) setupHealthCheck() {
	c.SystemContainer = NewSystemContainer(
		c.BaseRouter,
		c.Config,
	)
}

func (c *AppContainer) setupTransaction() {
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	errorHandler := errorhandler.NewErrorHandler(errorMapper, c.LoggersGroup.General)
	errorhandler.RegisterAllErrorHandlers(errorHandler)

	c.TransactionContainer = NewTransactionContainer(
		c.V1Router,
		c.DB,
		c.Config,
		errorHandler,
		c.LoggersGroup.Transaction,
	)
}

func (c *AppContainer) setupDetailing() {
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	errorHandler := errorhandler.NewErrorHandler(errorMapper, c.LoggersGroup.General)
	errorhandler.RegisterAllErrorHandlers(errorHandler)

	c.DetailingContainer = NewDetailingContainer(
		c.V1Router,
		c.DB,
		c.Config,
		errorHandler,
		c.LoggersGroup.Detailing,
	)
}
