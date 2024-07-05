package handlers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/waldrey/eulabs/internal/dto"
	"github.com/waldrey/eulabs/internal/infra/service"
	"github.com/waldrey/eulabs/tools"
)

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

// Get Product godoc
// @Summary      Get Product
// @Description  Get product by id
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200       {array}   entity.Product
// @Router       /api/v1/products/:id [get]
func (h *ProductHandler) FindOne(c echo.Context) error {
	log.Print("GET :id request initialization")

	id, err := tools.ValidateRequest(c)
	if err != nil {
		return err
	}

	product, err := h.Service.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	log.Print("GET :id request finished")
	return c.JSON(http.StatusOK, product)
}

// Delete Product godoc
// @Summary      Delete product
// @Description  Deletes the product from the system
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      204
// @Router       /api/v1/products/:id [delete]
func (h *ProductHandler) Delete(c echo.Context) error {
	log.Print("DELETE :id  request initialization")

	id, err := tools.ValidateRequest(c)
	if err != nil {
		return err
	}

	err = h.Service.Delete(id)
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	log.Print("DELETE :id request finished")
	return c.JSON(http.StatusNoContent, "")
}

// Update Product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200       {array}   entity.Product
// @Router       /api/v1/products/:id [put]
func (h *ProductHandler) UpdatePut(c echo.Context) error {
	log.Print("PUT :id  request initialization")

	id, err := tools.ValidateRequest(c)
	if err != nil {
		return err
	}

	var product dto.PutProductRequest
	if err := c.Bind(&product); err != nil {
		return err
	}

	if err := h.Validator.Struct(product); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": tools.FormatValidationError(err),
		})
	}

	_, err = h.Service.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	_, err = h.Service.Update(id, product)
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	productUpdated, err := h.Service.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	log.Print("PUT :id request finished")
	return c.JSON(http.StatusOK, productUpdated)
}

// Update Product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200       {array}   entity.Product
// @Router       /api/v1/products/:id [patch]
func (h *ProductHandler) UpdatePatch(c echo.Context) error {
	log.Print("PATCH :id  request initialization")

	id, err := tools.ValidateRequest(c)
	if err != nil {
		return err
	}

	var product dto.UpdateProductRequest
	if err := c.Bind(&product); err != nil {
		return err
	}

	if err := h.Validator.Struct(product); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": tools.FormatValidationError(err),
		})
	}

	_, err = h.Service.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	_, err = h.Service.Update(id, dto.PutProductRequest{
		Name:        tools.SafeDereferenceString(product.Name),
		Description: tools.SafeDereferenceString(product.Description),
		Price:       tools.SafeDereferenceFloat64(product.Price),
	})
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	productUpdated, err := h.Service.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}

	log.Print("PUT :id request finished")
	return c.JSON(http.StatusOK, productUpdated)
}
