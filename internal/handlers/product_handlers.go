package handlers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/waldrey/eulabs/internal/infra/service"
	"gorm.io/gorm"
)

type ProductDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	gorm.Model
}

type ProductHandler struct {
	Service   service.ProductInterface
	Validator *validator.Validate
}

func NewProductHandler(service service.ProductInterface) *ProductHandler {
	return &ProductHandler{
		Service:   service,
		Validator: validator.New(),
	}
}

// List Products godoc
// @Summary      List products
// @Description  Get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200       {array}   entity.Product
// @Router       /api/v1/products [get]
func (h *ProductHandler) List(c echo.Context) error {
	log.Print("GET request initialization")

	products, err := h.Service.FindAll()
	if err != nil {
		log.Print("Unknown error getting products in database")
		return c.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	log.Print("GET request finished")
	return c.JSON(http.StatusOK, products)
}
