package errormapper

import (
	"errors"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"gorm.io/gorm"
)

type GORMErrRecordNotFoundMapper struct{}

func (m *GORMErrRecordNotFoundMapper) MapError(err error) (error, bool) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrObjectNotFound, true
	}
	return nil, false
}

func BuildGORMErrorsMapperChain() *MapperChain {
	mc := NewChain()
	mc.registerMapper(&GORMErrRecordNotFoundMapper{})
	return mc
}
