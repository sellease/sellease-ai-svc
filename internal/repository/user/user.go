package user

import (
	"context"
	"sellease-ai/internal/entity/models"

	"gorm.io/gorm"
)

type RepositoryInterface interface {
	GetUserByPasscode(ctx context.Context, passCode string) (*models.User, error)
	GetUsers(ctx context.Context) (users []models.User, err error)
	AddUserImage(ctx context.Context, imageName, passcode string) (err error)
	GetUserImageById(ctx context.Context, passCode string) (imgName string, err error)
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) RepositoryInterface {
	return &userRepository{
		db: db,
	}
}
