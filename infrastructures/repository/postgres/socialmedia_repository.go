package repository

import (
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewSocialmediaRepository(database *gorm.DB) domains.SocialmediaRepository {
	return &SocialmediaRepositoryImpl{
		db: database,
	}
}

type SocialmediaRepositoryImpl struct {
	db *gorm.DB
}

func (socialmediarepo *SocialmediaRepositoryImpl) Insert(socialmedia entities.Socialmedia) error {
	err := socialmediarepo.db.Create(&socialmedia).Error
	if err != nil {
		return err
	}
	return nil
}