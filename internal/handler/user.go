package handler

import (
	"net/http"
	"sellease-ai/consts"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/usecase/user"
	"sellease-ai/internal/utils"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUserCase user.UsecaseInterface
}

func InitUserHandler(uc user.UsecaseInterface) *userHandler {
	return &userHandler{
		userUserCase: uc,
	}
}

func (h *userHandler) HandleUserLogin(c *gin.Context) {
	ctx := utils.GetContext(c)

	reqData, ok := c.Get(consts.TagUserLoginRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, utils.ErrInvalidRequest.Error()))
		return
	}

	loginRequest, ok := reqData.(request.UserLoginRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidParamsErrCode, utils.ErrInvalidParameter.Error()))
		return
	}

	user, err := h.userUserCase.GetUserByPasscode(ctx, loginRequest.Passcode)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidCredentialErrCode, err.Error()))
		return
	}

	token, err := h.userUserCase.GenerateJWTAccessToken(user.ID, user.Passcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(utils.AccessTokenErrCode, err.Error()))
		return
	}

	resp := response.UserLoginResponse{
		AccessToken: token,
	}

	c.JSON(http.StatusOK, utils.Send(resp))
}

func (h *userHandler) HandleAddUserImage(c *gin.Context) {
	ctx := utils.GetContext(c)

	reqData, ok := c.Get(consts.TagAddUserImageRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, utils.ErrInvalidRequest.Error()))
		return
	}

	request, ok := reqData.(request.AddUserImageRequest)
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.InvalidParamsErrCode, utils.ErrInvalidParameter.Error()))
		return
	}

	err := h.userUserCase.AddUserImage(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(utils.InternalServerErrCode, err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.Send(utils.UserImageAdded))
}

func (h *userHandler) HandleGetUserImage(c *gin.Context) {
	ctx, userId, err := fetchContextAndUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Fail(utils.ValidationErrCode, err.Error()))
		return
	}

	imgName, err := h.userUserCase.GetUserImageById(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(utils.InternalServerErrCode, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Send(imgName))
}
