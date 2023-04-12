package repository

import (
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewRoleRepository(db *gorm.DB) domains.RoleRepository{
	return &RoleRepositoryImpl{
		db: db,
	}
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (RoleRepo *RoleRepositoryImpl) Insert(role *entities.Role) error {
	err := RoleRepo.db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}
