package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	db *gorm.DB,
	cfg *config.Config,
	loggerGroup *logging.LoggerGroup,
) {
	v1Router := gin.Group("/api/v1/")

	registerTransactionRoutes(v1Router, db, cfg, *loggerGroup.Transaction)
}

func registerTransactionRoutes(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *config.Config,
	logger logging.Logger,
) {
	transactionsRouter := baseRouter.Group("/transactions/")
	transactionsRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	transactionRepository := repository.NewTransactionRepository(db)
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	transactionUsecase := usecase.NewTransactionUsecase(
		transactionRepository,
		errorMapper,
		cfg,
	)

	errorHandler := controller.DefaultErrorHandler()
	transactionController := controller.NewTransactionController(
		transactionUsecase,
		errorHandler,
		logger,
		cfg,
	)

	transactionRouter := NewTransactionRouter(transactionsRouter, transactionController, cfg)
	transactionRouter.RegisterAll()
}
