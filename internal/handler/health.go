package handler

import (
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/utils"

	"github.com/gin-gonic/gin"
)

type healthHandler struct{}

func InitHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) HandleGetHealth(c *gin.Context) {
	result := response.Health{
		Status:  "normal",
		Message: "system running normally",
	}
	c.JSON(200, utils.Send(result))
}
