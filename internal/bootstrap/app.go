package bootstrap

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/router"
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap/dependency"
	"github.com/FCTL3314/FinSight-transactions/internal/collections/slice"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Logger logging.Logger

	ginEngine *gin.Engine
	deps      *dependency.AppContainer
}

func NewApplication() *Application {
	deps := dependency.NewAppContainer()

	app := &Application{
		Logger:    deps.LoggersGroup.General,
		ginEngine: deps.GinEngine,
		deps:      deps,
	}
	app.initialize()

	return app
}

func (app *Application) initialize() {
	app.setGinMode()
	app.setGinTrustedProxies()
	app.registerGinRoutes()
}

func (app *Application) setGinMode() {
	modes := []string{gin.ReleaseMode, gin.DebugMode, gin.TestMode}

	if !slice.Contains(modes, app.deps.Config.Server.Mode) {
		app.Logger.Warn(
			"Unsupported Gin mode provided. Falling back to debug mode for safety.",
			logging.WithField("mode", app.deps.Config.Server.Mode),
			logging.WithField("allowed_modes", modes),
		)
		gin.SetMode(gin.DebugMode)
		return
	}

	gin.SetMode(app.deps.Config.Server.Mode)
}

func (app *Application) setGinTrustedProxies() {
	if err := app.ginEngine.SetTrustedProxies(app.deps.Config.Server.TrustedProxies); err != nil {
		app.Logger.Fatal(
			"Error setting trusted proxies",
			logging.WithError(err),
		)
	}
}

func (app *Application) registerGinRoutes() {
	router.RegisterAll(
		app.deps.TransactionContainer.RouterRegistrator,
		app.deps.SystemContainer.RouterRegistrator,
	)
}

func (app *Application) Run() error {
	addr := ":" + app.deps.Config.Server.Port

	app.Logger.Info("Listening and serving HTTP", logging.WithField("addr", addr))
	if err := app.ginEngine.Run(addr); err != nil {
		return err
	}
	return nil
}
