package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CORRELATION_HEADER = "x-correlation-id"
	CORRELATION_ID     = "correlation-id"
)

// Gin Middleware to set the correlation id in the gin context and response headers
func SetCorrelationIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// if the client has sent the correlation id, use that
		correlationId := ctx.GetHeader(CORRELATION_HEADER)
		if correlationId == "" { // if not, create new
			correlationId = uuid.New().String()
		}

		ctx.Set(CORRELATION_ID, correlationId)
		ctx.Header(CORRELATION_HEADER, correlationId)

		ctx.Next()
	}
}
