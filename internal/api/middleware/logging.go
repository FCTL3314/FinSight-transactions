package middleware

import (
	"github.com/FCTL3314/FinSight-transactions/internal/logging"
	"github.com/gin-gonic/gin"
)

func ErrorLoggerMiddleware(logger logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, ginErr := range c.Errors {
				fields := []logging.Field{
					logging.WithField("method", c.Request.Method),
					logging.WithField("path", c.Request.URL.Path),
					logging.WithField("error", ginErr.Error()),
					logging.WithField("type", ginErr.Type),
				}

				if ginErr.Meta != nil {
					if meta, ok := ginErr.Meta.(map[string]interface{}); ok {
						for key, value := range meta {
							fields = append(fields, logging.WithField(key, value))
						}
					}
				}

				logger.Error("Request processing error", fields...)
			}
		}

		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			logger.Error("HTTP error response",
				logging.WithField("method", c.Request.Method),
				logging.WithField("path", c.Request.URL.Path),
				logging.WithField("status_code", statusCode),
				logging.WithField("errors_count", len(c.Errors)),
			)
		}
	}
}
