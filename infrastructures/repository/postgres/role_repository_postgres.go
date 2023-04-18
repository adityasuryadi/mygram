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

func (RoleRepo *RoleRepositoryImpl) Edit(id string,role *entities.Role) error {
	var curRole *entities.Role
	tx:=RoleRepo.db.Begin()
	
	err := RoleRepo.db.Where("id = ?",id).First(&curRole).Error
	curRole.Name = role.Name
	tx.Save(&curRole) 
	if err != nil {
		return err
	}
	tx.Model(&curRole).Association("Permissions").Clear()
	if err != nil {
		tx.Rollback()
	}
	tx.Model(&curRole).Association("Permissions").Replace(role.Permissions)
	if err != nil {
		tx.Rollback()
	}
	if err := tx.Commit().Error; err !=  nil {
		tx.Rollback()
		// handle error
	}
	return nil
}
