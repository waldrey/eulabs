package database

import "github.com/waldrey/eulabs/internal/entity"

type ProductInterface interface {
	FindAll() ([]entity.Product, error)
}
