package handler

import (
	"context"
	"sellease-ai/consts"
	"sellease-ai/internal/utils"

	"github.com/gin-gonic/gin"
)

func fetchContextAndUserId(c *gin.Context) (context.Context, string, error) {
	ctx := utils.GetContext(c)
	data, ok := c.Get(consts.AuthUserIDKey)
	if !ok {
		return nil, "", utils.ErrInvalidUser
	}

	userId, ok := data.(string)
	if !ok {
		return nil, "", utils.ErrInvalidParameter
	}
	return ctx, userId, nil
}
