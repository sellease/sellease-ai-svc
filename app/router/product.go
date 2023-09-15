package router

import (
	"sellease-ai/consts"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/middleware"

	"github.com/gin-gonic/gin"
)

func (c *RouterContext) ProductRoutes(r *gin.RouterGroup) {
	productHandler := c.Handler.ProductHandler

	r.POST("/desc",
		middleware.Validate[request.ProductDescriptionRequest](consts.TagProductDescriptionRequest),
		productHandler.HandleGenerateProductDescription)
}
