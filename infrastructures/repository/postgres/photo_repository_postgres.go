package repository

import (
	"errors"
	"log"
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

// GetAll implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) GetAll() ([]*entities.Photo, error) {
	var photos []*entities.Photo
	err := repository.db.Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}

// InsertPhoto implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) InsertPhoto(photo entities.Photo) (*entities.Photo, error) {
	err := repository.db.Create(&photo).Error
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

// FindById implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) FindById(id string) (*entities.Photo, error) {
	var photo entities.Photo
	err := repository.db.Where("id = ?", id).First(&photo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &photo, nil
}

// UpdatePhoto implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) UpdatePhoto(id string, photo entities.Photo) error {
	err := repository.db.Save(&photo).Error
	if err != nil {
		return err
	}

	return nil
}

// DestroyPhoto implements domains.PhotoRepository
func (repository *PhotoRepositoryImpl) DestroyPhoto(id string) error {	
	err := repository.db.Where("id = ?",id).Delete(&entities.Photo{}).Error
	if err != nil {
		return err
	}
	log.Print(err)
	return nil
}
