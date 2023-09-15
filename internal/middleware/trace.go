package middleware

import (
	"sellease-ai/consts"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"

	"github.com/gin-gonic/gin"
)

// middleware that injects a 'X-Request-Id' into the context and header of each request
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		xRequestID := utils.GenerateUUID()
		c.Request.Header.Set(consts.XRequestIDKey, xRequestID)
		c.Set(consts.XRequestIDKey, xRequestID)
		c.Header(consts.XRequestIDKey, xRequestID)

		if c.Request.URL.Path != "/auth-engine/health" { // Ignoring health route from logging
			logger.Debugf(`[API-Hit] [X-Request-Id:%s] - "%s %s"`, xRequestID, c.Request.Method, c.Request.URL.Path)
		}
		c.Next()
	}
}
