package repository

import (
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewPermissionRepository(db *gorm.DB) domains.PermissionRepository{
	return &PermissionRepositoryImpl{
		db:db,
	}
}

type PermissionRepositoryImpl struct {
	db *gorm.DB
}

func (repository *PermissionRepositoryImpl) Insert(permission *entities.Permission) error {
	err := repository.db.Create(&repository).Error

	if err != nil {
		return err
	}
	return nil	
} 