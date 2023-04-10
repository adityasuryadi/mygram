package repository

import (
	"errors"
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

func (socialmediarepo *SocialmediaRepositoryImpl) GetAll()([]*entities.Socialmedia,error){
	var socialmedia []*entities.Socialmedia
	err:=socialmediarepo.db.Find(&socialmedia).Error
	if err != nil {
		return nil,err
	}
	return socialmedia,nil
}

func (socialmediarepo *SocialmediaRepositoryImpl) FindById(id string)(*entities.Socialmedia,error){
	var socialmedia entities.Socialmedia
	err := socialmediarepo.db.Where("id = ?", id).First(&socialmedia).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &socialmedia, nil
}


func (socialmediaRepo *SocialmediaRepositoryImpl) Update(socialmedia *entities.Socialmedia)(*entities.Socialmedia,error){
	err := socialmediaRepo.db.Save(&socialmedia).Error
	if err != nil {
		return nil,err
	}
	return socialmedia,nil
}


func(socialmediaRepo *SocialmediaRepositoryImpl) Delete(id string)error{
	var socialmedia entities.Socialmedia
	err:=socialmediaRepo.db.Where("id = ?", id).Delete(&socialmedia).Error
	if err != nil {
		return err
	}
	return nil
}