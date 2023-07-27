package mock_repository

import (
	"errors"
	entities "mygram/domains/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindById(id string) (*entities.Product,error) {
	arguments:=repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil,errors.New("Product Not Found")
	}else{
		product := arguments.Get(0).(entities.Product)
		return &product,nil
	}
}

func (repository *ProductRepositoryMock) Insert(product *entities.Product) error {
	panic("")
}



func (repository *ProductRepositoryMock) Update(product *entities.Product) (*entities.Product,error) {
	panic("")
}

func (repository *ProductRepositoryMock) Delete(id string) error {
	panic("")
}


