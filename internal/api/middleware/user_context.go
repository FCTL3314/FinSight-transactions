package middleware

import (
	"github.com/FCTL3314/FinSight-transactions/internal/api/controller"
	"github.com/FCTL3314/FinSight-transactions/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	UserIDHeader = "X-User-ID"
)

func UserContext(c *gin.Context) {
	userIDStr := c.Request.Header.Get(UserIDHeader)
	if userIDStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, domain.NewUnauthorizedErrorResponse())
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, domain.NewErrorResponse(
			"X-User-ID header must be a valid integer",
			"invalid_user_id",
		))
		return
	}

	c.Set(controller.UserIDContextKey, userID)
	c.Next()
}
