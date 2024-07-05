package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/waldrey/eulabs/internal/entity"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (p *ProductRepositoryMock) Create(product *entity.Product) (*entity.Product, error) {
	args := p.Called(product)
	if result, ok := args.Get(0).(*entity.Product); ok {
		result.ID = 1
		return result, args.Error(1)
	}
	return nil, args.Error(1)
}

func (p *ProductRepositoryMock) Delete(product *entity.Product) error {
	args := p.Called(product)
	return args.Error(0)
}

func (p *ProductRepositoryMock) FindAll() ([]entity.Product, error) {
	args := p.Called()
	if products, ok := args.Get(0).([]entity.Product); ok {
		return products, args.Error(1)
	}
	return nil, args.Error(1)
}

func (p *ProductRepositoryMock) FindByID(id int) (*entity.Product, error) {
	args := p.Called(id)
	if product, ok := args.Get(0).(*entity.Product); ok {
		return product, args.Error(1)
	}
	return nil, args.Error(1)
}

func (p *ProductRepositoryMock) Update(product *entity.Product) error {
	args := p.Called(product)
	return args.Error(0)
}
