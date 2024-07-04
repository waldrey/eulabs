package service

import (
	"github.com/waldrey/eulabs/internal/entity"
	"github.com/waldrey/eulabs/internal/infra/database"
)

type Product struct {
	repository database.ProductInterface
}

func ProductService(repository database.ProductInterface) *Product {
	return &Product{repository: repository}
}

func (p *Product) FindAll() ([]entity.Product, error) {
	return p.repository.FindAll()
}
