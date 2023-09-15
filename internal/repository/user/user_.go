package user

import (
	"context"
	"fmt"
	"sellease-ai/config"
	"sellease-ai/internal/entity/models"
	"sellease-ai/logger"

	"gorm.io/gorm"
)

func (r *userRepository) GetUserByPasscode(ctx context.Context, passCode string) (user *models.User, err error) {
	schema := config.GetConfig().PostgresSchema
	query := fmt.Sprintf(readUserByPasscodeQuery, schema)

	err = r.db.Raw(query, passCode).Scan(&user).Error
	if err != nil {
		logger.WithContext(ctx).Errorf("error fetching from user - passcode: %s - %s ", passCode, err.Error())
		return user, err
	}
	if user == nil {
		logger.WithContext(ctx).Errorf("error fetching user - passcode:%s - %s ", passCode, gorm.ErrRecordNotFound)
		return user, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context) (users []models.User, err error) {
	schema := config.GetConfig().PostgresSchema
	query := fmt.Sprintf(readUsersQuery, schema)

	err = r.db.Raw(query).Scan(&users).Error
	if err != nil {
		logger.WithContext(ctx).Errorf("error fetching users - %s ", err.Error())
		return users, err
	}

	return users, nil
}

func (r *userRepository) AddUserImage(ctx context.Context, imageName, passcode string) (err error) {
	schema := config.GetConfig().PostgresSchema
	query := fmt.Sprintf(addUserImageQuery, schema)

	err = r.db.Exec(
		query,
		imageName,
		passcode,
	).Error
	if err != nil {
		logger.WithContext(ctx).Errorf("error adding user image [passcode:%s] - %s",
			passcode, err.Error())
		return err
	}

	return nil
}
func (r *userRepository) GetUserImageById(ctx context.Context, userId string) (imgName string, err error) {
	schema := config.GetConfig().PostgresSchema
	query := fmt.Sprintf(readImageNameByIdQuery, schema)

	err = r.db.Raw(query, userId).Scan(&imgName).Error
	if err != nil {
		logger.WithContext(ctx).Errorf("error fetching from user - user_id: %s - %s ", userId, err.Error())
		return imgName, err
	}
	if len(imgName) < 1 {
		logger.WithContext(ctx).Errorf("error fetching image_name - user_id:%s - %s ", userId, gorm.ErrRecordNotFound)
		return imgName, gorm.ErrRecordNotFound
	}

	return imgName, nil
}
