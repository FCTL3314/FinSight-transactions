package errormapper

import (
	"errors"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"gorm.io/gorm"
)

func MapGORMRecordNotFoundError(err error) (error, bool) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrObjectNotFound, true
	}
	return nil, false
}

func BuildGORMErrorsMapperChain() MapperChain {
	mc := NewMapperChain()
	mc.registerMapper(MapGORMRecordNotFoundError)
	return mc
}
