package database

import (
	"github.com/waldrey/eulabs/internal/entity"
)

type ProductInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindByID(id int) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(product *entity.Product) error
}
