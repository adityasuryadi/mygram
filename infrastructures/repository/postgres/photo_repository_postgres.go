package repository

import (
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewPhotoRepository(database *gorm.DB) domains.PhotoRepository {
	return &PhotoRepositoryImpl{
		db: database,
	}
}

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

// InsertPhoto implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) InsertPhoto(photo entities.Photo) (*entities.Photo,error){
	err := repository.db.Create(&photo).Error
	if err != nil {
		return nil,err
	}
	return &photo,nil
}
