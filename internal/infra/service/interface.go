package service

import "github.com/waldrey/eulabs/internal/entity"

type ProductInterface interface {
	FindAll() ([]entity.Product, error)
}
