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

func (repository *UserTokenRepositoryImpl) InsertToken(user *entities.User, token string) {
	userToken := &entities.UserToken{
		UserId:    user.Id,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 72),
	}
	repository.db.Create(userToken)
}

func (repository *UserTokenRepositoryImpl) RemoveToken() {
}
