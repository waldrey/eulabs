package service

import (
	"github.com/waldrey/eulabs/internal/dto"
	"github.com/waldrey/eulabs/internal/entity"
)

type ProductInterface interface {
	FindAll() ([]entity.Product, error)
	FindOne(id int) (*entity.Product, error)
	Update(id int, product dto.PutProductRequest) (*entity.Product, error)
	Delete(id int) error
}
