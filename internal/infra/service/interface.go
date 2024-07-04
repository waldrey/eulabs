package service

import "github.com/waldrey/eulabs/internal/entity"

type ProductInterface interface {
	FindAll() ([]entity.Product, error)
	FindOne(id int) (*entity.Product, error)
	Delete(id int) error
}
