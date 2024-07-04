package handlers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/waldrey/eulabs/internal/infra/database"
	"gorm.io/gorm"
)

type ProductDTO struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	gorm.Model
}

type ProductHandler struct {
	Repository database.ProductInterface
	Validator  *validator.Validate
}

func NewProductHandler(repository database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		Repository: repository,
		Validator:  validator.New(),
	}
}

func (h *ProductHandler) List(c echo.Context) error {
	log.Print("GET request initialization")

	products, err := h.Repository.FindAll()
	if err != nil {
		log.Print("Unknown error getting products in database")
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	log.Print("GET request finished")
	return c.JSON(http.StatusOK, products)
}
