package errorhandler

import (
	"net/http"

	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/FCTL3314/FinSight-transactions/internal/errormapper"
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	errorMapper errormapper.MapperChain
	handlers    []HandlerFunc
	Logger      logging.Logger
}

func NewErrorHandler(
	errorMapper errormapper.MapperChain,
	logger logging.Logger,
) *ErrorHandler {
	return &ErrorHandler{
		errorMapper: errorMapper,
		Logger:      logger,
	}
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

	eh.Logger.Error(
		"Unhandled error occurred",
		logging.WithError(err),
	)

	c.JSON(http.StatusInternalServerError, domain.InternalServerErrorResponse)
}
