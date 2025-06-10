package bootstrap

import (
	"fmt"
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap/container"
	"github.com/FCTL3314/FinSight-transactions/internal/collections"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Router      *gin.RouterGroup
	Config      *config.Config
	LoggerGroup *logging.LoggerGroup

	ginEngine *gin.Engine
	container *container.AppContainer
}

func NewApplication() *Application {
	c := container.NewAppContainer()

	app := &Application{
		Router:      c.Router,
		Config:      c.Config,
		LoggerGroup: c.LoggerGroup,
		ginEngine:   c.GinEngine,
		container:   c,
	}
	app.initialize()

	return app
}

func (app *Application) Run() error {
	addr := ":" + app.Config.Server.Port

	app.LoggerGroup.General.Info(fmt.Sprintf("Listening and serving HTTP on %s\n", addr))
	if err := app.ginEngine.Run(addr); err != nil {
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

	if !collections.Contains(modes, app.Config.Server.Mode) {
		app.LoggerGroup.General.Warn(
			"Unsupported Gin mode provided. Falling back to DebugMode for safety.",
			logging.WithField("mode", app.Config.Server.Mode),
			logging.WithField("allowed_modes", modes),
		)
		gin.SetMode(gin.DebugMode)
		return
	}

	gin.SetMode(app.Config.Server.Mode)
}

func (app *Application) setGinTrustedProxies() {
	if err := app.ginEngine.SetTrustedProxies(app.Config.Server.TrustedProxies); err != nil {
		app.LoggerGroup.General.Fatal(
			"Error setting trusted proxies",
			logging.WithError(err),
		)
	}
}

func (app *Application) registerGinRoutes() {
	router.RegisterAll(
		app.container.Transaction.RouterRegistrator,
	)
}
