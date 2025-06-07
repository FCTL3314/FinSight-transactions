package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type TransactionRouter struct {
	router                *gin.RouterGroup
	transactionController controller.TransactionController
	cfg                   *config.Config
}

func NewTransactionRouter(
	router *gin.RouterGroup,
	transactionController controller.TransactionController,
	cfg *config.Config,
) *TransactionRouter {
	return &TransactionRouter{router, transactionController, cfg}
}

func (wr *TransactionRouter) RegisterAll() {
	wr.RegisterGet()
	wr.RegisterList()
	wr.RegisterCreate()
	wr.RegisterUpdate()
	wr.RegisterDelete()
}

func (wr *TransactionRouter) RegisterGet() {
	wr.router.GET("/:id", wr.transactionController.Get)
}

func (wr *TransactionRouter) RegisterList() {
	wr.router.GET("", wr.transactionController.List)
}

func (wr *TransactionRouter) RegisterCreate() {
	wr.router.POST("", wr.transactionController.Create)
}

func (wr *TransactionRouter) RegisterUpdate() {
	wr.router.PATCH("/:id", wr.transactionController.Update)
}

func (wr *TransactionRouter) RegisterDelete() {
	wr.router.DELETE("/:id", wr.transactionController.Delete)
}
