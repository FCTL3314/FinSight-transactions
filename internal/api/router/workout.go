package router

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/gin-gonic/gin"
)

type WorkoutRouter struct {
	router            *gin.RouterGroup
	workoutController controller.WorkoutController
	cfg               *config.Config
}

func NewWorkoutRouter(
	router *gin.RouterGroup,
	workoutController *controller.DefaultWorkoutController,
	cfg *config.Config,
) *WorkoutRouter {
	return &WorkoutRouter{router, workoutController, cfg}
}

func (wr *WorkoutRouter) RegisterAll() {
	wr.RegisterGet()
	wr.RegisterList()
	wr.RegisterCreate()
	wr.RegisterUpdate()
	wr.RegisterDelete()
}

func (wr *WorkoutRouter) RegisterGet() {
	wr.router.GET("/:id", wr.workoutController.Get)
}

func (wr *WorkoutRouter) RegisterList() {
	wr.router.GET("", wr.workoutController.List)
}

func (wr *WorkoutRouter) RegisterCreate() {
	wr.router.POST("", wr.workoutController.Create)
}

func (wr *WorkoutRouter) RegisterUpdate() {
	wr.router.PATCH("/:id", wr.workoutController.Update)
}

func (wr *WorkoutRouter) RegisterDelete() {
	wr.router.DELETE("/:id", wr.workoutController.Delete)
}
