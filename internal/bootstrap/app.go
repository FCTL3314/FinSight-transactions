package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Application struct {
	Router *gin.Engine
	DB     *gorm.DB
	Cfg    *Config
}

func NewApplication() *Application {
	var app Application
	app.initConfig()
	app.initDB()
	app.initGin()
	return &app
}

func (app *Application) initConfig() {
	c, err := Load()

	if err != nil {
		log.Fatal("Error during config loading. Please check if environmental files exist.")
	}
	app.Cfg = c
}

func (app *Application) initDB() {
	DBConnector := NewGormConnector(
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
	switch app.Cfg.Server.GinMode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
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

func (app *Application) Run() {
	if err := app.Router.Run(":" + app.Cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
