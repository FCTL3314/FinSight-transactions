package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type TransactionRouter interface {
	Router
}

type transactionRouter struct {
	router                *gin.RouterGroup
	transactionController controller.TransactionController
	cfg                   *config.Config
}

func NewTransactionRouter(
	baseRouter *gin.RouterGroup,
	transactionController controller.TransactionController,
	cfg *config.Config,
) TransactionRouter {
	baseRoute := baseRouter.Group("/transactions/")
	return &transactionRouter{baseRoute, transactionController, cfg}
}

func (tr *transactionRouter) Get() {
	tr.router.GET("/:id", middleware.UserContext, tr.transactionController.Get)
}

func (tr *transactionRouter) List() {
	tr.router.GET("", middleware.UserContext, tr.transactionController.List)
}

func (tr *transactionRouter) Create() {
	tr.router.POST("", middleware.UserContext, tr.transactionController.Create)
}

func (tr *transactionRouter) Update() {
	tr.router.PATCH("/:id", middleware.UserContext, tr.transactionController.Update)
}

func (tr *transactionRouter) Delete() {
	tr.router.DELETE("/:id", middleware.UserContext, tr.transactionController.Delete)
}

type transactionRouterRegistrator struct {
	router TransactionRouter
}

func NewTransactionRouterRegistrator(transactionRouter TransactionRouter) Registrator {
	return &transactionRouterRegistrator{router: transactionRouter}
}

func (r *transactionRouterRegistrator) Register() {
	r.router.Get()
	r.router.List()
	r.router.Create()
	r.router.Update()
	r.router.Delete()
}
