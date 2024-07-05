package entity

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidName        = errors.New("invalid name")
	ErrInvalidDescription = errors.New("invalid description")
	ErrInvalidPrice       = errors.New("invalid price")
)

type Product struct {
	gorm.Model  `swaggerignore:"true"`
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := product.IsValid()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) IsValid() error {
	if p.Name == "" {
		return ErrInvalidName
	}

	if p.Description == "" {
		return ErrInvalidDescription
	}

	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
