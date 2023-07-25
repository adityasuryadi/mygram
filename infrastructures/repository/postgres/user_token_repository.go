package repository

import (
	"time"

	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewUserTokenRepository(database *gorm.DB) domains.UserTokenRepository {
	return &UserTokenRepositoryImpl{
		db: database,
	}
}

type UserTokenRepositoryImpl struct {
	db *gorm.DB
}

// InsertTokenInsertTokenWithTx implements domains.UserTokenRepository
func (repository *UserTokenRepositoryImpl) InsertTokenWithTx(tx *gorm.DB, user *entities.User, token string) error {
	userToken := &entities.UserToken{
		UserId:    user.Id,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 72),
	}

	err := tx.Create(userToken).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserTokenRepositoryImpl) RemoveToken(userId string) error {
	var userToken entities.UserToken
	err := repository.db.Unscoped().Where("user_id = ?", userId).Delete(&userToken).Error

	if err != nil {
		return err
	}
	return nil
}
