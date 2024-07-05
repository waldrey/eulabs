package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waldrey/eulabs/internal/dto"
	"github.com/waldrey/eulabs/internal/entity"
	"github.com/waldrey/eulabs/test/mock"
)

func TestGivenAValidParams_WhenICallProductCreateService_ThenShouldReceiveProductWithAllParams(t *testing.T) {
	productRequest := dto.CreateProductRequest{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	productEntityService := &entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	repository := &mock.ProductRepositoryMock{}
	repository.On("Create", productEntityService).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil)
	service := ProductService(repository)

	product, err := service.Create(productRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro", product.Name)
	assert.Equal(t, "O poderoso computador da Apple", product.Description)
	assert.Equal(t, 23000.00, product.Price)

	repository.AssertExpectations(t)
}

func TestGivenAValidParams_WhenICallProductDeleteService_ThenShouldReceiveStatusNotContent(t *testing.T) {
	productRequest := dto.CreateProductRequest{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	productEntityService := &entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	repository := &mock.ProductRepositoryMock{}
	repository.On("Create", productEntityService).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil)
	repository.On("FindByID", 1).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil)
	repository.On("Delete", &entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}).Return(nil)
	service := ProductService(repository)

	product, err := service.Create(productRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro", product.Name)
	assert.Equal(t, "O poderoso computador da Apple", product.Description)
	assert.Equal(t, 23000.00, product.Price)

	err = service.Delete(int(product.ID))
	assert.NoError(t, err)

	repository.AssertExpectations(t)
}

func TestGivenAValidParams_WhenICallUpdateProductService_ThenShouldReceiveSuccess(t *testing.T) {
	productRequest := dto.CreateProductRequest{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	productEntityService := &entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	repository := &mock.ProductRepositoryMock{}
	repository.On("Create", productEntityService).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil)
	repository.On("FindByID", 1).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil).Once()
	repository.On("Update", &entity.Product{
		Name:        "Macbook Pro 2024",
		Description: "O poderoso computador da Apple",
		Price:       15000.00,
	}).Return(nil)
	repository.On("FindByID", 1).Return(&entity.Product{
		Name:        "Macbook Pro 2024",
		Description: "O poderoso computador da Apple",
		Price:       15000.00,
	}, nil).Once()
	service := ProductService(repository)

	product, err := service.Create(productRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro", product.Name)
	assert.Equal(t, "O poderoso computador da Apple", product.Description)
	assert.Equal(t, 23000.00, product.Price)

	product, err = service.Update(int(product.ID), dto.PutProductRequest{
		Name:        "Macbook Pro 2024",
		Description: "O poderoso computador da Apple",
		Price:       15000.00,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro 2024", product.Name)
	assert.Equal(t, 15000.00, product.Price)

	repository.AssertExpectations(t)
}

func TestGivenAValidParams_WhenICallFindAllProductService_ThenShouldReceiveSuccess(t *testing.T) {
	repository := &mock.ProductRepositoryMock{}
	repository.On("FindAll").Return([]entity.Product{
		{Id: 1, Name: "Macbook Pro", Description: "Description", Price: 100.0},
		{Id: 2, Name: "iPhone 15 Pro Max", Description: "Description", Price: 50.60},
		{Id: 2, Name: "Livro Domain-Driven Design", Description: "Description", Price: 1599.99},
	}, nil)
	service := ProductService(repository)

	expectedProducts := []entity.Product{
		{Id: 1, Name: "Macbook Pro", Description: "Description", Price: 100.0},
		{Id: 2, Name: "iPhone 15 Pro Max", Description: "Description", Price: 50.60},
		{Id: 2, Name: "Livro Domain-Driven Design", Description: "Description", Price: 1599.99},
	}

	products, err := service.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, products)
	repository.AssertExpectations(t)
}

func TestGivenAValidParams_WhenICallFindOneProductService_ThenShouldReceiveSuccess(t *testing.T) {
	productRequest := dto.CreateProductRequest{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	productEntityService := &entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}

	repository := &mock.ProductRepositoryMock{}
	repository.On("Create", productEntityService).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil)
	repository.On("FindByID", 1).Return(&entity.Product{
		Name:        "Macbook Pro",
		Description: "O poderoso computador da Apple",
		Price:       23000.00,
	}, nil).Once()
	service := ProductService(repository)

	product, err := service.Create(productRequest)
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro", product.Name)
	assert.Equal(t, "O poderoso computador da Apple", product.Description)
	assert.Equal(t, 23000.00, product.Price)

	product, err = service.FindOne(int(product.ID))
	assert.NoError(t, err)
	assert.Equal(t, "Macbook Pro", product.Name)
	repository.AssertExpectations(t)
}
