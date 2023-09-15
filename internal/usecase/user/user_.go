package user

import (
	"context"
	"errors"

	"sellease-ai/config"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"

	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func (u *userUsecase) GenerateJWTAccessToken(userId, passcode string) (string, error) {
	// prepare claims for token
	claims := &models.JWTData{
		StandardClaims: jwt.StandardClaims{
			// set token lifetime in timestamp
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		// add custom claims like user_id or email,
		// it can vary according to requirements
		CustomClaims: map[string]string{
			"user_id":  userId,
			"passcode": passcode,
		},
	}

	// generate a string using claims and HS256 algorithm
	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the generated key using secretKey
	secretKey := config.GetConfig().JWTSecret
	token, err := tokenString.SignedString([]byte(secretKey))
	if err != nil {
		logger.Errorf("error generating token - %s", err.Error())
	}

	return token, err
}

func (u *userUsecase) GetUserByPasscode(ctx context.Context, passCode string) (*models.User, error) {
	result, err := u.userRepo.GetUserByPasscode(ctx, passCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, utils.ErrUserNotFound
		}
		return result, utils.ErrDBFailedToFetchData
	}

	return result, nil
}

func (u *userUsecase) AddUserImage(ctx context.Context, req request.AddUserImageRequest) (err error) {
	err = u.userRepo.AddUserImage(ctx, req.ImageName, req.Passcode)
	if err != nil {
		return utils.ErrDBFailedToUpdateData
	}
	return nil
}

func (u *userUsecase) GetUserImageById(ctx context.Context, userId string) (imgName string, err error) {
	imgName, err = u.userRepo.GetUserImageById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imgName, utils.ErrUserNotFound
		}
		return imgName, utils.ErrDBFailedToFetchData
	}

	return imgName, nil
}
