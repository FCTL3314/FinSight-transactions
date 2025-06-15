package errormapper

type MapperFunc func(err error) (mappedErr error, ok bool)

type MapperChain interface {
	registerMapper(mapper MapperFunc)
	getMappers() []MapperFunc
	MapError(err error) error
}

type mapperChain struct {
	mappers []MapperFunc
}

func NewMapperChain() MapperChain {
	return &mapperChain{}
}

func (mc *mapperChain) registerMapper(mapper MapperFunc) {
	mc.mappers = append(mc.mappers, mapper)
}

func (mc *mapperChain) getMappers() []MapperFunc {
	return mc.mappers
}

func (mc *mapperChain) MapError(err error) error {
	for _, mapper := range mc.mappers {
		if mappedErr, ok := mapper(err); ok {
			return mappedErr
		}
	}
	return err
}

func BuildAllErrorsMapperChain() MapperChain {
	mc := NewMapperChain()
	GORMMapperChain := BuildGORMErrorsMapperChain()
	PostgresMapperChain := BuildPostgresErrorsMapperChain()

	allMapperChains := [2]MapperChain{
		GORMMapperChain,
		PostgresMapperChain,
	}

	for _, mapperChain := range allMapperChains {

		for _, mapper := range mapperChain.getMappers() {
			mc.registerMapper(mapper)
		}
	}

	return mc
}
