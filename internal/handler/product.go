package handler

import (
	"net/http"
	"sellease-ai/consts"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/usecase/product"
	"sellease-ai/internal/utils"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUseCase product.UsecaseInterface
}

func InitProductHandler(uc product.UsecaseInterface) *productHandler {
	return &productHandler{
		productUseCase: uc,
	}
}

func (h *productHandler) HandleGenerateProductDescription(c *gin.Context) {
	ctx := utils.GetContext(c)

	reqData, ok := c.Get(consts.TagProductDescriptionRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, utils.ErrInvalidRequest.Error()))
		return
	}

	prodRequest, ok := reqData.(request.ProductDescriptionRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidParamsErrCode, utils.ErrInvalidParameter.Error()))
		return
	}

	productDescription, err := h.productUseCase.GenerateProductDesc(ctx, prodRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidCredentialErrCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Send(productDescription))
}

func (h *productHandler) HandleGenerateKeywords(c *gin.Context) {
	ctx := utils.GetContext(c)

	query := c.Query("query")
	result, err := h.productUseCase.GenerateKeywords(ctx, query)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidCredentialErrCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Send(result))
}

func (h *productHandler) HandleTranslateText(c *gin.Context) {
	ctx := utils.GetContext(c)

	reqData, ok := c.Get(consts.TagTranslationRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, utils.ErrInvalidRequest.Error()))
		return
	}

	request, ok := reqData.(request.TranslationRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidParamsErrCode, utils.ErrInvalidParameter.Error()))
		return
	}

	resp, err := h.productUseCase.TranslateText(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(utils.InternalServerErrCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Send(resp))
}
