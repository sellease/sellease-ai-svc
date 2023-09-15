package utils

import (
	"context"
	"sellease-ai/consts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetContext(c *gin.Context) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, consts.XRequestID(consts.XRequestIDKey), GetXRequestId(c))
	return ctx
}

// GetXRequestId returns 'X-Request-Id'
func GetXRequestId(c *gin.Context) string {
	xRequestId := getRequestIDFromHeaders(c)
	if len(xRequestId) > 0 {
		return xRequestId
	}
	return getRequestIDFromContext(c)
}

// getRequestIDFromHeaders returns 'X-Request-Id' from the headers if present
func getRequestIDFromHeaders(c *gin.Context) string {
	return c.Request.Header.Get(consts.XRequestIDKey)
}

// getRequestIDFromContext returns 'X-Request-Id' from the given context if present
func getRequestIDFromContext(c *gin.Context) string {
	if v, ok := c.Get(consts.XRequestIDKey); ok {
		if requestID, ok := v.(string); ok {
			return requestID
		}
	}
	return ""
}

// generate uuid as string
func GenerateUUID() string {
	return uuid.New().String()
}
