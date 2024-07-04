package service

import (
	"log"

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

func (p *Product) FindOne(id int) (*entity.Product, error) {
	return p.repository.FindByID(id)
}

func (p *Product) Delete(id int) error {
	product, err := p.repository.FindByID(id)
	if err != nil {
		return err
	}

	log.Print("record found to deletion")
	return p.repository.Delete(product)
}
