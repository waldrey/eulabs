package handlers

import (
	"log"
	"net/http"
	"strconv"

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

// Delete Product godoc
// @Summary      Delete product
// @Description  Deletes the product from the system
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      204	[]
// @Router       /api/v1/products/:id [delete]
func (h *ProductHandler) Delete(c echo.Context) error {
	log.Print("DELETE request initialization")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID must be an integer"})
	}

	if id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID must be a positive integer"})
	}

	err = h.Service.Delete(id)
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	log.Print("DELETE request finished")
	return c.JSON(http.StatusNoContent, "")
}
