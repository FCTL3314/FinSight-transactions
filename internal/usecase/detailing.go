package usecase

import (
	"github.com/FCTL3314/FinSight-transactions/internal/config"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/repository"
	"github.com/FCTL3314/FinSight-transactions/pkg/schemas"
)

type DetailingUsecase interface {
	Get(authUserId, id int64) (*domain.FinanceDetailing, error)
	List(authUserId int64, params *domain.Params) (*domain.PaginatedResult[*domain.FinanceDetailing], error)
	Create(authUserId int64, createFinanceDetailingRequest *schemas.CreateFinanceDetailingRequest) (*domain.FinanceDetailing, error)
	Update(authUserId, id int64, updateFinanceDetailingRequest *schemas.UpdateFinanceDetailingRequest) (*domain.FinanceDetailing, error)
	Delete(authUserId, id int64) error
}

type detailingUsecase struct {
	detailingRepository repository.DetailingRepository
	cfg                 *config.Config
}

func NewDetailingUsecase(
	detailingRepository repository.DetailingRepository,
	cfg *config.Config,
) DetailingUsecase {
	return &detailingUsecase{
		detailingRepository: detailingRepository,
		cfg:                 cfg,
	}
}

func (du *detailingUsecase) Get(authUserId, id int64) (*domain.FinanceDetailing, error) {
	filterParams := &domain.FilterParams{
		Conditions: []domain.FilterCondition{
			{Field: "id", Operator: domain.OpEq, Value: id},
			{Field: "user_id", Operator: domain.OpEq, Value: authUserId},
		},
	}
	return du.detailingRepository.Get(filterParams)
}

func (du *detailingUsecase) List(authUserId int64, params *domain.Params) (*domain.PaginatedResult[*domain.FinanceDetailing], error) {
	params.Filter.Conditions = append(params.Filter.Conditions, domain.FilterCondition{
		Field:    "user_id",
		Operator: domain.OpEq,
		Value:    authUserId,
	})

	detailings, err := du.detailingRepository.Fetch(params)
	if err != nil {
		return nil, err
	}

	count, err := du.detailingRepository.Count(params.Filter)
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedResult[*domain.FinanceDetailing]{Results: detailings, Count: count}, nil
}

func (du *detailingUsecase) Create(authUserId int64, createFinanceDetailingRequest *schemas.CreateFinanceDetailingRequest) (*domain.FinanceDetailing, error) {
	detailing := createFinanceDetailingRequest.ToDomainModel(authUserId)

	filterParams := domain.NewFilterParams(
		domain.FilterCondition{
			Field:    "user_id",
			Operator: domain.OpEq,
			Value:    authUserId,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpGte,
			Value:    createFinanceDetailingRequest.DateFrom,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpLte,
			Value:    createFinanceDetailingRequest.DateTo,
		},
	)

	return du.detailingRepository.Create(detailing, filterParams)
}

func (du *detailingUsecase) Update(authUserId, id int64, updateFinanceDetailingRequest *schemas.UpdateFinanceDetailingRequest) (*domain.FinanceDetailing, error) {
	detailingToUpdate, err := du.Get(authUserId, id)
	if err != nil {
		return nil, err
	}

	updateFinanceDetailingRequest.ApplyToDomainModel(detailingToUpdate)

	filterParams := domain.NewFilterParams(
		domain.FilterCondition{
			Field:    "user_id",
			Operator: domain.OpEq,
			Value:    authUserId,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpGte,
			Value:    detailingToUpdate.DateFrom,
		},
		domain.FilterCondition{
			Field:    "made_at",
			Operator: domain.OpLte,
			Value:    detailingToUpdate.DateTo,
		},
	)

	return du.detailingRepository.Update(detailingToUpdate, filterParams)
}

func (du *detailingUsecase) Delete(authUserId, id int64) error {
	detailing, err := du.Get(authUserId, id)
	if err != nil {
		return err
	}

	if detailing.UserID != authUserId {
		return domain.ErrAccessDenied
	}

	return du.detailingRepository.Delete(id)
}
