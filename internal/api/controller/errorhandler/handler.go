package errorhandler

import (
	"errors"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc = func(c *gin.Context, err error) bool

func handleObjectNotFound(c *gin.Context, err error) bool {
	if errors.Is(err, domain.ErrObjectNotFound) {
		c.JSON(http.StatusNotFound, domain.NotFoundResponse)
		return true
	}
	return false
}

func handleAccessDenied(c *gin.Context, err error) bool {
	if errors.Is(err, domain.ErrAccessDenied) {
		c.JSON(http.StatusForbidden, domain.ForbiddenResponse)
		return true
	}
	return false
}

func handleInvalidParam(c *gin.Context, err error) bool {
	var errInvalidURLParam *domain.ErrInvalidURLParam
	if errors.As(err, &errInvalidURLParam) {
		c.JSON(http.StatusBadRequest, domain.NewInvalidURLParamResponse(errInvalidURLParam.Error()))
		return true
	}
	return false
}

func handlePaginationLimitExceeded(c *gin.Context, err error) bool {
	var errPaginationLimit *domain.ErrPaginationLimitExceeded
	if errors.As(err, &errPaginationLimit) {
		c.JSON(http.StatusBadRequest, domain.NewPaginationErrorResponse(errPaginationLimit.Error()))
		return true
	}
	return false
}

func handleUniqueConstraint(c *gin.Context, err error) bool {
	var errObjectUniqueConstraint *domain.ErrObjectUniqueConstraint
	if errors.As(err, &errObjectUniqueConstraint) {
		c.JSON(http.StatusConflict, domain.NewUniqueConstraintErrorResponse(err.Error()))
		return true
	}
	return false
}
