package repository

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/pkg/models"
	"gorm.io/gorm"
)

type SourceOfIncomeRepository interface {
	Repository[models.SourceOfIncome]
}

type DefaultSourceOfIncomeRepository struct {
	db        *gorm.DB
	toPreload []string
}

func NewDefaultSourceOfIncomeRepository(db *gorm.DB) *DefaultSourceOfIncomeRepository {
	return &DefaultSourceOfIncomeRepository{db: db}
}

func (wr *DefaultSourceOfIncomeRepository) GetById(id int64) (*models.SourceOfIncome, error) {
	return wr.Get(&domain.FilterParams{
		Query: "id = ?",
		Args:  []interface{}{id},
	})
}

func (wr *DefaultSourceOfIncomeRepository) Get(filterParams *domain.FilterParams) (*models.SourceOfIncome, error) {
	var sourceOfIncome models.SourceOfIncome
	query := wr.db.Where(filterParams.Query, filterParams.Args...)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if err := (query.First(&sourceOfIncome)).Error; err != nil {
		return nil, err
	}

	return &sourceOfIncome, nil
}

func (wr *DefaultSourceOfIncomeRepository) Fetch(params *domain.Params) ([]*models.SourceOfIncome, error) {
	var sourcesOfIncome []*models.SourceOfIncome
	query := wr.db.Where(params.Filter.Query, params.Filter.Args...)
	query = query.Order(params.Order)
	query = applyPreloadsForGORMQuery(query, wr.toPreload)
	if params.Pagination.Limit != 0 {
		query = query.Limit(params.Pagination.Limit).Offset(params.Pagination.Offset)
	}
	if err := (query.Find(&sourcesOfIncome)).Error; err != nil {
		return nil, err
	}

	return sourcesOfIncome, nil
}

func (wr *DefaultSourceOfIncomeRepository) Create(sourceOfIncome *models.SourceOfIncome) (*models.SourceOfIncome, error) {
	if err := (wr.db.Save(&sourceOfIncome)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.SourceOfIncome{}), wr.toPreload)
	if err := query.First(sourceOfIncome).Error; err != nil {
		return nil, err
	}

	return sourceOfIncome, nil
}

func (wr *DefaultSourceOfIncomeRepository) Update(sourceOfIncome *models.SourceOfIncome) (*models.SourceOfIncome, error) {
	if err := (wr.db.Save(&sourceOfIncome)).Error; err != nil {
		return nil, err
	}

	query := applyPreloadsForGORMQuery(wr.db.Model(&models.SourceOfIncome{}), wr.toPreload)
	if err := query.First(sourceOfIncome).Error; err != nil {
		return nil, err
	}

	return sourceOfIncome, nil
}

func (wr *DefaultSourceOfIncomeRepository) Delete(id int64) error {
	result := wr.db.Where("id = ?", id).Delete(&models.SourceOfIncome{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (wr *DefaultSourceOfIncomeRepository) Count(params *domain.FilterParams) (int64, error) {
	var count int64
	if err := (wr.db.Model(&models.SourceOfIncome{}).Where(params.Query, params.Args...).Count(&count)).Error; err != nil {
		return 0, err
	}
	return count, nil
}
