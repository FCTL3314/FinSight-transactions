package bootstrap

import (
	"fmt"
	"github.com/FCTL3314/FinSight-transactions/internal/collections"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"

	// "github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Router      *gin.Engine
	DB          *gorm.DB
	Cfg         *config.Config
	LoggerGroup *logging.LoggerGroup
}

func NewApplication() *Application {
	var app Application
	app.initConfig()
	app.initDB()
	app.initGin()
	app.initLoggerGroup()
	return &app
}

func (app *Application) Run() {
	// router.RegisterRoutes(app.Router, app.DB, app.Cfg, app.LoggerGroup)

	addr := ":" + app.Cfg.Server.Port

	fmt.Printf("Listening and serving HTTP on %s\n", addr)
	if err := app.Router.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func (app *Application) initConfig() {
	c, err := config.Load()
	if err != nil {
		log.Fatal("Error during config loading. Please check if environmental files exist.")
	}
	app.Cfg = c
}

func (app *Application) initDB() {
	DBConnector := NewConnector(
		app.Cfg.Database.Name,
		app.Cfg.Database.User,
		app.Cfg.Database.Password,
		app.Cfg.Database.Host,
		app.Cfg.Database.Port,
	)
	db, err := DBConnector.Connect()
	if err != nil {
		log.Fatal("Error during database connection.")
	}
	app.DB = db
}

func (app *Application) setGinMode() {
	modes := []string{gin.ReleaseMode, gin.DebugMode, gin.TestMode}

	if collections.Contains(modes, app.Cfg.Server.Mode) {
		gin.SetMode(app.Cfg.Server.Mode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func (app *Application) initGin() {
	app.setGinMode()

	r := gin.Default()
	if err := r.SetTrustedProxies(app.Cfg.Server.TrustedProxies); err != nil {
		log.Fatal(err)
	}

	app.Router = r
}

func (app *Application) initLoggerGroup() {
	transactionLogger := logging.InitTransactionLogger()

	loggerGroup := logging.NewLoggerGroup(
		transactionLogger,
	)
	app.LoggerGroup = loggerGroup
}
