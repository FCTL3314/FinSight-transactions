package repository

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"gorm.io/gorm"
)

type IWorkoutRepository interface {
	domain.Repository[models.Workout]
}

type WorkoutRepository struct {
	db        *gorm.DB
	toPreload []string
}

func NewWorkoutRepository(db *gorm.DB) *WorkoutRepository {
	return &WorkoutRepository{db: db, toPreload: []string{"User", "WorkoutExercises", "WorkoutExercises.Exercise"}}
}

func (wr *WorkoutRepository) GetById(id int64) (*models.Workout, error) {
	return wr.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (wr *WorkoutRepository) Get(filterParams *domain.FilterParams) (*models.Workout, error) {
	var workout models.Workout
	query := wr.db.Where(filterParams.Query, filterParams.Args...)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if err := (query.First(&workout)).Error; err != nil {
		return nil, err
	}

	return &workout, nil
}

func (wr *WorkoutRepository) Fetch(params *domain.Params) ([]*models.Workout, error) {
	var workouts []*models.Workout
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&workouts)).Error; err != nil {
		return nil, err
	}

	return workouts, nil
}

func (wr *WorkoutRepository) Create(workout *models.Workout) (*models.Workout, error) {
	if err := (wr.db.Save(&workout)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.Workout{}), wr.toPreload)
	if err := query.First(workout).Error; err != nil {
		return nil, err
	}

	return workout, nil
}

func (wr *WorkoutRepository) Update(workout *models.Workout) (*models.Workout, error) {
	if err := (wr.db.Omit("WorkoutExercises").Save(&workout)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.Workout{}), wr.toPreload)
	if err := query.First(workout).Error; err != nil {
		return nil, err
	}

	return workout, nil
}

func (wr *WorkoutRepository) Delete(id int64) error {
	result := wr.db.Where("id = ?", id).Delete(&models.Workout{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (wr *WorkoutRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (wr.db.Model(&models.Workout{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}

type WorkoutExerciseRepository struct {
	db        *gorm.DB
	toPreload []string
}

func NewWorkoutExerciseRepository(db *gorm.DB) *WorkoutExerciseRepository {
	return &WorkoutExerciseRepository{db: db, toPreload: []string{"Workout", "Exercise"}}
}

func (wr *WorkoutExerciseRepository) GetById(id int64) (*models.WorkoutExercise, error) {
	return wr.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (wr *WorkoutExerciseRepository) Get(filterParams *domain.FilterParams) (*models.WorkoutExercise, error) {
	var workout models.WorkoutExercise
	query := wr.db.Where(filterParams.Query, filterParams.Args...)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if err := (query.First(&workout)).Error; err != nil {
		return nil, err
	}

	return &workout, nil
}

// Fetch TODO: Здесь можно вместо получения WorkoutExercise как вложенные объекты, сделать отдельный endpoint, тогда будет легче добавить возможность сортировки.
func (wr *WorkoutExerciseRepository) Fetch(params *domain.Params) ([]*models.WorkoutExercise, error) {
	var workouts []*models.WorkoutExercise
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&workouts)).Error; err != nil {
		return nil, err
	}

	return workouts, nil
}

func (wr *WorkoutExerciseRepository) Create(workoutExercise *models.WorkoutExercise) (*models.WorkoutExercise, error) {
	if err := (wr.db.Save(&workoutExercise)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.WorkoutExercise{}), wr.toPreload)
	if err := query.First(workoutExercise).Error; err != nil {
		return nil, err
	}

	return workoutExercise, nil
}

func (wr *WorkoutExerciseRepository) Update(workoutExercise *models.WorkoutExercise) (*models.WorkoutExercise, error) {
	if err := (wr.db.Save(&workoutExercise)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.WorkoutExercise{}), wr.toPreload)
	if err := query.First(workoutExercise).Error; err != nil {
		return nil, err
	}

	return workoutExercise, nil
}

func (wr *WorkoutExerciseRepository) Delete(id int64) error {
	result := wr.db.Where("id = ?", id).Delete(&models.WorkoutExercise{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (wr *WorkoutExerciseRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (wr.db.Model(&models.WorkoutExercise{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
