package middleware

import (
	"context"
	"errors"
	"net/http"

	"sellease-ai/config"
	"sellease-ai/consts"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/usecase/user"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"

	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	BEARER        = "Bearer "
	MISSING_AUTH  = "MISSING_AUTH"
	INVALID_TOKEN = "INVALID_TOKEN"
	INACTIVE_USER = "INACTIVE_USER"

	// Header
	Authorization = "Authorization"

	// Json message keys
	StatusCodeKey = "code"
	MessageKey    = "message"
	DetailsKey    = "details"
)

func Authorize(u user.UsecaseInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := utils.GetContext(ctx)
		authToken := ctx.GetHeader(Authorization)
		if len(authToken) < 1 {
			abortRequest(ctx, MISSING_AUTH, utils.ErrMissingAuthorization)
		}
		splitToken := strings.Split(authToken, BEARER)
		accessToken := splitToken[1]
		claims, err := ParseJWTAccessToken(accessToken)
		if err != nil {
			abortRequest(ctx, INVALID_TOKEN, err)
			return
		}
		if !checkUserIsActive(c, u, claims.CustomClaims["passcode"]) {
			abortRequest(ctx, INACTIVE_USER, utils.ErrInActiveUser)
			return
		}
		ctx.Set(consts.AuthUserIDKey, claims.CustomClaims["user_id"])
	}
}

func ParseJWTAccessToken(accessToken string) (*models.JWTData, error) {
	secretKey := config.GetConfig().JWTSecret
	token, err := jwt.ParseWithClaims(accessToken, &models.JWTData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		logger.Errorf("error in parse jwt token - %s", err.Error())
		return nil, err
	}

	claims := token.Claims.(*models.JWTData)
	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("access token expired")
	}

	return claims, err
}

func abortRequest(c *gin.Context, abortType string, err error) {
	switch abortType {
	case MISSING_AUTH:
		c.JSON(http.StatusUnauthorized, gin.H{
			StatusCodeKey: http.StatusUnauthorized,
			MessageKey:    MISSING_AUTH,
			DetailsKey:    utils.ErrMissingAuthorization.Error(),
		})
	case INVALID_TOKEN:
		c.JSON(http.StatusUnauthorized, gin.H{
			StatusCodeKey: http.StatusUnauthorized,
			MessageKey:    INVALID_TOKEN,
			DetailsKey:    err.Error(),
		})
	case INACTIVE_USER:
		c.JSON(http.StatusUnauthorized, gin.H{
			StatusCodeKey: http.StatusUnauthorized,
			MessageKey:    INACTIVE_USER,
			DetailsKey:    err.Error(),
		})
		c.Abort()
	}
}

func checkUserIsActive(ctx context.Context, u user.UsecaseInterface, passCode string) bool {
	_, err := u.GetUserByPasscode(ctx, passCode)
	return err == nil
}
