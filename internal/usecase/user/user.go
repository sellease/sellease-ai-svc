package user

import (
	"context"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/repository/user"
)

type UsecaseInterface interface {
	// User
	GetUserByPasscode(ctx context.Context, passCode string) (*models.User, error)
	AddUserImage(ctx context.Context, req request.AddUserImageRequest) (err error)
	GetUserImageById(ctx context.Context, userId string) (imgName string, err error)

	// Access Token
	GenerateJWTAccessToken(userId, passcode string) (string, error)
}

type userUsecase struct {
	userRepo user.RepositoryInterface
}

func InitUserUsecase(
	user user.RepositoryInterface,
) UsecaseInterface {
	return &userUsecase{
		userRepo: user,
	}
}
