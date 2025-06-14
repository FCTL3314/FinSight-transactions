package errorhandler

import (
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorHandler struct {
	errorMapper errormapper.MapperChain
	handlers    []HandlerFunc
}

func NewErrorHandler(errorMapper errormapper.MapperChain) *ErrorHandler {
	return &ErrorHandler{errorMapper: errorMapper}
}

func (eh *ErrorHandler) RegisterHandler(handler HandlerFunc) {
	eh.handlers = append(eh.handlers, handler)
}

func (eh *ErrorHandler) Handle(c *gin.Context, err error) {
	mappedErr := err
	if eh.errorMapper != nil {
		mappedErr = eh.errorMapper.MapError(err)
	}

	for _, handler := range eh.handlers {
		if handler(c, mappedErr) {
			return
		}
	}
	c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
}
