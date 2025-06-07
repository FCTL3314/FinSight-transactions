package middleware

import (
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorLoggerMiddleware(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := c.Writer.Status()

		if statusCode != http.StatusInternalServerError {
			return
		}

		logger.Error("Internal server error",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
		)
	}
}
