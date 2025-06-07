package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/api/middleware"
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
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
	loggerGroup *bootstrap.LoggerGroup,
) {
	v1Router := gin.Group("/api/v1/")

	registerWorkoutRoutes(v1Router, db, cfg, *loggerGroup.Workout)
}

func registerWorkoutRoutes(
	baseRouter *gin.RouterGroup,
	db *gorm.DB,
	cfg *config.Config,
	logger bootstrap.Logger,
) {
	workoutsRouter := baseRouter.Group("/workouts/")
	workoutsRouter.Use(middleware.ErrorLoggerMiddleware(logger))

	workoutRepository := repository.NewWorkoutRepository(db)
	workoutExerciseRepository := repository.NewWorkoutExerciseRepository(db)
	errorMapper := errormapper.BuildAllErrorsMapperChain()
	workoutUsecase := usecase.NewWorkoutUsecase(
		workoutRepository,
		*workoutExerciseRepository,
		errorMapper,
		cfg,
	)

	errorHandler := controller.DefaultErrorHandler()
	workoutController := controller.NewWorkoutController(
		workoutUsecase,
		errorHandler,
		logger,
		cfg,
	)

	workoutRouter := NewWorkoutRouter(workoutsRouter, workoutController, cfg)
	workoutRouter.RegisterAll()
}
