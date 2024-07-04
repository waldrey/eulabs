package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyPrice_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{
		ID:          12345,
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
	}
	assert.Error(t, product.IsValid(), "invalid price")
}

func TestGivenAnEmptyName_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{
		ID:          12345,
		Description: "O poderoso computador da Apple",
	}
	assert.Error(t, product.IsValid(), "invalid name")
}

func TestGivenAnEmptyDescription_WhenCreateANewProduct_ThenShouldReceiveAnError(t *testing.T) {
	product := Product{
		ID: 12345,
	}
	assert.Error(t, product.IsValid(), "invalid description")
}

func TestGivenAnEmptyDescription_WhenICallNewProductFunc_ThenShouldReceiveAnError(t *testing.T) {
	_, err := NewProduct("Lego McLaren F1", "", 1400.00)
	assert.Error(t, err, "invalid description")
}

func TestGivenAnEmptyName_WhenICallNewProductFunc_ThenShouldReceiveAnError(t *testing.T) {
	_, err := NewProduct("", "A poderosa McLaren F1 de 2022", 1400.00)
	assert.Error(t, err, "invalid name")
}

func TestGivenAValidParams_WhenICallNewProduct_ThenShouldReceiveCreateProductWithAllParams(t *testing.T) {
	product := Product{
		ID:          1234,
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}
	assert.Equal(t, 1234, product.ID)
	assert.Equal(t, 23000.0, product.Price)
	assert.Equal(t, "O poderoso computador da Apple", product.Description)
	assert.Nil(t, product.IsValid())
}

func TestGivenAValidParams_WhenICallNewProductFunc_ThenShouldReceiveCreateProductWithAllParams(t *testing.T) {
	product, err := NewProduct("Lego McLaren F1", "A poderosa McLaren F1 de 2022", 1400.00)
	assert.Nil(t, err)
	assert.Equal(t, "Lego McLaren F1", product.Name)
	assert.Equal(t, "A poderosa McLaren F1 de 2022", product.Description)
	assert.Equal(t, 1400.00, product.Price)
}
