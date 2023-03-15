package repository

import (
	domains "mygram/domains/user"
	entities "mygram/domains/user/entity"

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
func (repository *UserRepositoryImpl) GetUserByEmail() {
	panic("unimplemented")
}
