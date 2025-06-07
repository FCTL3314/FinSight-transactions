package bootstrap

import (
	"fmt"
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap/container"
	"github.com/FCTL3314/FinSight-transactions/internal/collections"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Router      *gin.Engine
	DB          *gorm.DB
	Config      *config.Config
	LoggerGroup *logging.LoggerGroup

	container *container.AppContainer
}

func NewApplication() *Application {
	c := container.NewAppContainer()

	app := &Application{
		Router:      c.Router,
		DB:          c.DB,
		Config:      c.Config,
		LoggerGroup: c.LoggerGroup,
		container:   c,
	}

	app.initialize()

	return app
}

func (app *Application) Run() error {
	addr := ":" + app.Config.Server.Port

	app.LoggerGroup.General.Info(fmt.Sprintf("Listening and serving HTTP on %s\n", addr))
	if err := app.Router.Run(addr); err != nil {
		return err
	}
	return nil
}

func (app *Application) initialize() {
	app.setGinMode()
	app.setGinTrustedProxies()
	app.registerGinRoutes()
}

func (app *Application) setGinMode() {
	modes := []string{gin.ReleaseMode, gin.DebugMode, gin.TestMode}

	if collections.Contains(modes, app.Config.Server.Mode) {
		gin.SetMode(app.Config.Server.Mode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}

func (app *Application) setGinTrustedProxies() {
	if err := app.Router.SetTrustedProxies(app.Config.Server.TrustedProxies); err != nil {
		app.LoggerGroup.General.Fatal(
			"error setting trusted proxies",
			logging.WithError(err),
		)
	}
}

func (app *Application) registerGinRoutes() {
	router.RegisterRoutes(app.Router, app.container)
}
