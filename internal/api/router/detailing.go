package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type DetailingRouter interface {
	GetRouter
}

type detailingRouter struct {
	router              *gin.RouterGroup
	detailingController controller.DetailingController
	cfg                 *config.Config
}

func NewDetailingRouter(
	baseRouter *gin.RouterGroup,
	detailingController controller.DetailingController,
	cfg *config.Config,
) DetailingRouter {
	baseRoute := baseRouter.Group("/detailing/")
	return &detailingRouter{baseRoute, detailingController, cfg}
}

func (tr *detailingRouter) Get() {
	tr.router.GET("/", middleware.UserContext, tr.detailingController.Get)
}

type detailingRouterRegistrator struct {
	router DetailingRouter
}

func NewDetailingRouterRegistrator(detailingRouter DetailingRouter) Registrator {
	return &detailingRouterRegistrator{router: detailingRouter}
}

func (r *detailingRouterRegistrator) Register() {
	r.router.Get()
}
