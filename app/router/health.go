package router

import "github.com/gin-gonic/gin"

func (c *RouterContext) HealthRoutes(r *gin.RouterGroup) {
	healthHandler := c.Handler.HealthHandler

	r.GET("/health", healthHandler.HandleGetHealth)
}
