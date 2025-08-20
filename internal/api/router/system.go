package router

import (
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type SystemRouter interface {
	GetRouter
}

type systemRouter struct {
	router *gin.RouterGroup
	cfg    *config.Config
}

func NewSystemRouter(
	baseRouter *gin.RouterGroup,
	cfg *config.Config,
) SystemRouter {
	return &systemRouter{baseRouter, cfg}
}

func (tr *systemRouter) Get() {
	tr.router.GET("/health-check", func(c *gin.Context) { c.Status(http.StatusOK) })
}

type systemRouterRegistrator struct {
	router SystemRouter
}

func NewSystemRouterRegistrator(systemRouter SystemRouter) Registrator {
	return &systemRouterRegistrator{router: systemRouter}
}

func (r *systemRouterRegistrator) Register() {
	r.router.Get()
}
