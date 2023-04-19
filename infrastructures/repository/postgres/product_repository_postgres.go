package repository

import (
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
