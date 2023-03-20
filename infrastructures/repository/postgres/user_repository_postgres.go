package repository

import (
	"errors"
	domains "mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewUserRepositoryPostgres(database *gorm.DB) domains.UserRepository {
	return &UserRepositoryImpl{
		db: database,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Insert implements domains.UserRepository
func (repository *UserRepositoryImpl) Insert(user entities.User) error{
	result := repository.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByEmail implements domains.UserRepository
func (repository *UserRepositoryImpl) GetUserByEmail(email string) (entities.User,error) {
	var userEntity entities.User
	err := repository.db.Where("email",email).First(&userEntity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.User{},err
	}
	return userEntity,nil	
}
