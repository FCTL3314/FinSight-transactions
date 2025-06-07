package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap/container"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type GetRouter interface {
	RegisterGet()
}

type ListRouter interface {
	RegisterList()
}

type CreateRouter interface {
	RegisterCreate()
}

type UpdateRouter interface {
	RegisterUpdate()
}

type DeleteRouter interface {
	RegisterDelete()
}

type AllRouter interface {
	RegisterAll()
}

type Router interface {
	GetRouter
	ListRouter
	CreateRouter
	UpdateRouter
	DeleteRouter
	AllRouter
}

func RegisterRoutes(
	gin *gin.Engine,
	container *container.AppContainer,
) {
	v1Router := gin.Group("/api/v1/")

	registerTransactionRoutes(
		v1Router,
		container.Transaction.Controller,
		container.Config,
		container.LoggerGroup.Transaction,
	)
}

func registerTransactionRoutes(
	baseRouter *gin.RouterGroup,
	transactionController controller.TransactionController,
	cfg *config.Config,
	logger logging.Logger,
) {
	transactionsRouter := baseRouter.Group("/transactions/")
	transactionsRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	transactionRouter := NewTransactionRouter(
		transactionsRouter,
		transactionController,
		cfg,
	)
	transactionRouter.RegisterAll()
}
