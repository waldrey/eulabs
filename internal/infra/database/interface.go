package database

import "github.com/waldrey/eulabs/internal/entity"

type ProductInterface interface {
	FindAll() ([]entity.Product, error)
	FindByID(id int) (*entity.Product, error)
	Delete(product *entity.Product) error
}
