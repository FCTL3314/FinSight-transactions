package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type TransactionRouter interface {
	Router
}

type DefaultTransactionRouter struct {
	router                *gin.RouterGroup
	transactionController controller.TransactionController
	cfg                   *config.Config
}

func NewDefaultTransactionRouter(
	baseRouter *gin.RouterGroup,
	transactionController controller.TransactionController,
	cfg *config.Config,
	logger logging.Logger,
) *DefaultTransactionRouter {
	baseRoute := baseRouter.Group("/transactions/")
	baseRoute.Use(middleware.ErrorLoggerMiddleware(logger))

	return &DefaultTransactionRouter{baseRoute, transactionController, cfg}
}

func (tr *DefaultTransactionRouter) Get() {
	tr.router.GET("/:id", tr.transactionController.Get)
}

func (tr *DefaultTransactionRouter) List() {
	tr.router.GET("", tr.transactionController.List)
}

func (tr *DefaultTransactionRouter) Create() {
	tr.router.POST("", tr.transactionController.Create)
}

func (tr *DefaultTransactionRouter) Update() {
	tr.router.PATCH("/:id", tr.transactionController.Update)
}

func (tr *DefaultTransactionRouter) Delete() {
	tr.router.DELETE("/:id", tr.transactionController.Delete)
}

type TransactionRouterRegistrator struct {
	router TransactionRouter
}

func NewTransactionRouterRegistrator(transactionRouter TransactionRouter) *TransactionRouterRegistrator {
	return &TransactionRouterRegistrator{router: transactionRouter}
}

func (r *TransactionRouterRegistrator) Register() {
	r.router.Get()
	r.router.List()
	r.router.Create()
	r.router.Update()
	r.router.Delete()
}
