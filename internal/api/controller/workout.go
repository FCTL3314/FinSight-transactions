package controller

import (
	"github.com/FCTL3314/FinSight-transactions/internal/bootstrap"
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	usecase2 "github.com/FCTL3314/FinSight-transactions/internal/usecase"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkoutController interface {
	Controller
}

type DefaultWorkoutController struct {
	usecase      usecase2.IWorkoutUsecase
	errorHandler *ErrorHandler
	Logger       bootstrap.Logger
	cfg          *config.Config
}

func NewWorkoutController(
	usecase usecase2.IWorkoutUsecase,
	errorHandler *ErrorHandler,
	logger bootstrap.Logger,
	cfg *config.Config,
) *DefaultWorkoutController {
	return &DefaultWorkoutController{
		usecase:      usecase,
		errorHandler: errorHandler,
		Logger:       logger,
		cfg:          cfg,
	}
}

func (wc *DefaultWorkoutController) Get(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	workout, err := wc.usecase.GetById(id)

	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := workout.ToResponseWorkout()

	c.JSON(http.StatusOK, responseWorkout)
}

func (wc *DefaultWorkoutController) List(c *gin.Context) {
	params, err := getParams(c, wc.cfg.Pagination.MaxWorkoutLimit)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	paginatedResult, err := wc.usecase.List(&params)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkouts := models.ToResponseWorkouts(paginatedResult.Results)

	paginatedResponse := domain.PaginatedResponse{
		Count:   paginatedResult.Count,
		Limit:   params.Pagination.Limit,
		Offset:  params.Pagination.Offset,
		Results: responseWorkouts,
	}

	c.JSON(http.StatusOK, paginatedResponse)
}

func (wc *DefaultWorkoutController) Create(c *gin.Context) {
	var workout models.CreateWorkoutRequest
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	createdWorkout, err := wc.usecase.Create(authUserId, &workout)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := createdWorkout.ToResponseWorkout()

	c.JSON(http.StatusCreated, responseWorkout)
}

func (wc *DefaultWorkoutController) Update(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	var workout models.UpdateWorkoutRequest
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, domain.NewValidationErrorResponse(err.Error()))
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	updatedWorkout, err := wc.usecase.Update(authUserId, id, &workout)
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	responseWorkout := updatedWorkout.ToResponseWorkout()

	c.JSON(http.StatusOK, responseWorkout)
}

func (wc *DefaultWorkoutController) Delete(c *gin.Context) {
	id, err := getParamAsInt64(c, "id")
	if err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}

	authUserId := c.GetInt64(string(UserIDContextKey))

	if err := wc.usecase.Delete(authUserId, id); err != nil {
		wc.errorHandler.Handle(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
