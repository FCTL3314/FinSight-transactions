package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type DetailingRouter interface {
	GetRouter
	ListRouter
	CreateRouter
	UpdateRouter
	DeleteRouter
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
	baseRoute := baseRouter.Group("/detailing")
	return &detailingRouter{baseRoute, detailingController, cfg}
}

func (tr *detailingRouter) Get() {
	tr.router.GET("/:id", middleware.UserContext, tr.detailingController.Get)
}

func (tr *detailingRouter) List() {
	tr.router.GET("", middleware.UserContext, tr.detailingController.List)
}

func (tr *detailingRouter) Create() {
	tr.router.POST("/", middleware.UserContext, tr.detailingController.Create)
}

func (tr *detailingRouter) Update() {
	tr.router.PATCH("/:id", middleware.UserContext, tr.detailingController.Update)
}

func (tr *detailingRouter) Delete() {
	tr.router.DELETE("/:id", middleware.UserContext, tr.detailingController.Delete)
}

type detailingRouterRegistrator struct {
	router DetailingRouter
}

func NewDetailingRouterRegistrator(detailingRouter DetailingRouter) Registrator {
	return &detailingRouterRegistrator{router: detailingRouter}
}

func (r *detailingRouterRegistrator) Register() {
	r.router.Get()
	r.router.List()
	r.router.Create()
	r.router.Update()
	r.router.Delete()
}
