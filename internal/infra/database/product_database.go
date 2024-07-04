package database

import (
	"github.com/waldrey/eulabs/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func ProductRepository(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := p.DB.Find(&products).Error

	return products, err
}

func (p *Product) Delete(product *entity.Product) error {
	err := p.DB.Delete(product).Error
	return err
}

func (p *Product) FindByID(id int) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}
