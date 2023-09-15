package handler

import (
	"net/http"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/usecase/fileproc"
	"sellease-ai/internal/utils"

	"github.com/gin-gonic/gin"
)

type fileProcHandler struct {
	fileProcUseCase fileproc.UsecaseInterface
}

func InitFileProcHandler(uc fileproc.UsecaseInterface) *fileProcHandler {
	return &fileProcHandler{
		fileProcUseCase: uc,
	}
}

func (h *fileProcHandler) HandleProcessFile(c *gin.Context) {
	ctx := utils.GetContext(c)

	var req request.FileUploadRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, utils.ErrInvalidRequest.Error()))
		return
	}

	resp, err := h.fileProcUseCase.ProcessFile(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(utils.InternalServerErrCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Send(resp))
}
