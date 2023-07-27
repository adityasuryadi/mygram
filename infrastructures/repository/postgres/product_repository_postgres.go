package repository

import (
	"errors"
	"mygram/domains"
	entities "mygram/domains/entity"

	"gorm.io/gorm"
)

func NewProductRepository(db *gorm.DB) domains.ProductRepository{
	return &ProductRepositoryImpl{
		db: db,
	}
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func (repository *ProductRepositoryImpl) Insert(product *entities.Product) error {
	err := repository.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}


func (repository *ProductRepositoryImpl) FindById(id string) (*entities.Product,error) {
	var product entities.Product
	err := repository.db.Where("id = ? ",id).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &product,nil
}

func (repository *ProductRepositoryImpl) Update(product *entities.Product) (*entities.Product,error) {
	err := repository.db.Save(&product).Error
	if err != nil {
		return nil,err
	}
	return product,nil
}

func (repository *ProductRepositoryImpl) Delete(id string) error {
	var product entities.Product
	err := repository.db.Where("id = ?",id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

