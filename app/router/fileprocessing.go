package router

import (
	"github.com/gin-gonic/gin"
)

func (c *RouterContext) FileProcessingRoutes(r *gin.RouterGroup) {
	fileProcHandler := c.Handler.FileProcHandler

	r.POST("/process-csv",
		fileProcHandler.HandleProcessFile)
}
