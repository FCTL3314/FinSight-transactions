package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
)

type IWorkoutUsecase interface {
	GetById(id int64) (*models.Workout, error)
	Get(params *domain.FilterParams) (*models.Workout, error)
	List(params *domain.Params) (*domain.PaginatedResult[*models.Workout], error)
	Create(authUserId int64, createWorkoutRequest *models.CreateWorkoutRequest) (*models.Workout, error)
	Update(authUserId, id int64, updateWorkoutRequest *models.UpdateWorkoutRequest) (*models.Workout, error)
	Delete(authUserId, id int64) error
}

type WorkoutUsecase struct {
	workoutRepository         repository.IWorkoutRepository
	workoutExerciseRepository repository.WorkoutExerciseRepository
	errorMapper               errormapper.Chain
	cfg                       *config.Config
}

func NewWorkoutUsecase(
	workoutRepository repository.IWorkoutRepository,
	workoutExerciseRepository repository.WorkoutExerciseRepository,
	errorMapper errormapper.Chain,
	cfg *config.Config,
) *WorkoutUsecase {
	return &WorkoutUsecase{
		workoutRepository:         workoutRepository,
		workoutExerciseRepository: workoutExerciseRepository,
		errorMapper:               errorMapper,
		cfg:                       cfg,
	}
}

func (wu *WorkoutUsecase) GetById(id int64) (*models.Workout, error) {
	return &models.Workout{}, nil
}

func (wu *WorkoutUsecase) Get(params *domain.FilterParams) (*models.Workout, error) {
	return &models.Workout{}, nil
}

func (wu *WorkoutUsecase) List(params *domain.Params) (*domain.PaginatedResult[*models.Workout], error) {
	return &domain.PaginatedResult[*models.Workout]{}, nil
}

func (wu *WorkoutUsecase) Create(authUserId int64, createWorkoutRequest *models.CreateWorkoutRequest) (*models.Workout, error) {
	return &models.Workout{}, nil
}

func (wu *WorkoutUsecase) Update(authUserId int64, id int64, updateWorkoutRequest *models.UpdateWorkoutRequest) (*models.Workout, error) {
	return &models.Workout{}, nil
}

func (wu *WorkoutUsecase) Delete(authUserId int64, id int64) error {
	return nil
}
