package handlers

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/waldrey/eulabs/internal/dto"
	"github.com/waldrey/eulabs/internal/entity"
	"github.com/waldrey/eulabs/internal/infra/service"
	"github.com/waldrey/eulabs/pkg/requests"
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

// Create Product godoc
// @Summary      Create Product
// @Description  Create product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      201       {array}   requests.TypeSuccessResponse
// @Failure 	 400 	   {object}  requests.TypeErrorResponse
// @Failure 	 422 	   {object}  requests.TypeErrorResponse
// @Failure 	 500 	   {object}  requests.TypeErrorResponse
// @Router       /api/v1/products [post]
func (h *ProductHandler) Create(c echo.Context) error {
	log.Print("POST request initialization")

	var product dto.CreateProductRequest
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": tools.FormatValidationError(err),
		})
	}

	if err := h.Validator.Struct(product); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": tools.FormatValidationError(err),
		})
	}

	err := h.Service.Create(product)
	if err != nil {
		errResponse := requests.ErrorResponse("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	log.Print("POST request finished")
	successResponse := requests.SuccessResponse(entity.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
	return c.JSON(http.StatusCreated, successResponse)
}

// List Products godoc
// @Summary      List products
// @Description  Get all products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200       {array}   requests.TypeSuccessResponse
// @Failure 	 500 	   {object}  requests.TypeErrorResponse
// @Router       /products [get]
func (h *ProductHandler) List(c echo.Context) error {
	log.Print("GET request initialization")

	products, err := h.Service.FindAll()
	if err != nil {
		log.Print("Unknown error getting products in database")
		errResponse := requests.ErrorResponse("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	log.Print("GET request finished")
	successResponse := requests.SuccessListResponse(products)
	return c.JSON(http.StatusOK, successResponse)
}

// Get Product godoc
// @Summary      Get Product
// @Description  Get product by id
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(int)
// @Success      200       {array}   requests.TypeSuccessResponse
// @Failure      400       {object}  requests.TypeErrorResponse
// @Failure 	 404 	   {object}  requests.TypeErrorResponse
// @Failure      500       {object}  requests.TypeErrorResponse
// @Router       /products/{id} [get]
func (h *ProductHandler) FindOne(c echo.Context) error {
	log.Print("GET :id request initialization")

	id, err := tools.ValidateRequest(c)
	if err != nil {
		return err
	}

	product, err := h.Service.FindOne(id)
	if err != nil {
		errResponse := requests.ErrorResponse("Product not found")
		return c.JSON(http.StatusNotFound, errResponse)
	}

	log.Print("GET :id request finished")
	successResponse := requests.SuccessResponse(*product)
	return c.JSON(http.StatusOK, successResponse)
}

// Delete Product godoc
// @Summary      Delete product
// @Description  Deletes the product from the system
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(int)
// @Success      204
// @Failure      400       {object}  requests.TypeErrorResponse
// @Failure      404       {object}  requests.TypeErrorResponse
// @Failure      500       {object}  requests.TypeErrorResponse
// @Router       /products/{id} [delete]
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
			errResponse := requests.ErrorResponse("Product not found")
			return c.JSON(http.StatusNotFound, errResponse)
		}
		errResponse := requests.ErrorResponse("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	log.Print("DELETE :id request finished")
	return c.JSON(http.StatusNoContent, "")
}

// Update Product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(int)
// @Param        request     body      dto.PutProductRequest  true  "product request"
// @Success      200       {array}   requests.TypeSuccessResponse
// @Failure      400       {object}  requests.TypeErrorResponse
// @Failure      404       {object}  requests.TypeErrorResponse
// @Failure      422       {object}  requests.TypeErrorResponse
// @Failure      500       {object}  requests.TypeErrorResponse
// @Router       /products/{id} [put]
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
		errResponse := requests.ErrorResponse("Product not found")
		return c.JSON(http.StatusNotFound, errResponse)
	}

	_, err = h.Service.Update(id, product)
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			errResponse := requests.ErrorResponse("Product not found")
			return c.JSON(http.StatusNotFound, errResponse)
		}
		errResponse := requests.ErrorResponse("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	productUpdated, err := h.Service.FindOne(id)
	if err != nil {
		errResponse := requests.ErrorResponse("Product not found")
		return c.JSON(http.StatusNotFound, errResponse)
	}

	log.Print("PUT :id request finished")
	successResponse := requests.SuccessResponse(*productUpdated)
	return c.JSON(http.StatusOK, successResponse)
}

// Update Product godoc
// @Summary      Update product
// @Description  Update product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(int)
// @Param        request     body      dto.UpdateProductRequest  true  "product request"
// @Success      200       {array}   requests.TypeSuccessResponse
// @Failure      422       {object}  requests.TypeErrorResponse
// @Failure      404       {object}  requests.TypeErrorResponse
// @Failure      500       {object}  requests.TypeErrorResponse
// @Router       /products/{id} [patch]
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
		errResponse := requests.ErrorResponse("Product not found")
		return c.JSON(http.StatusNotFound, errResponse)
	}

	_, err = h.Service.Update(id, dto.PutProductRequest{
		Name:        tools.SafeDereferenceString(product.Name),
		Description: tools.SafeDereferenceString(product.Description),
		Price:       tools.SafeDereferenceFloat64(product.Price),
	})
	if err != nil {
		log.Print("Unknown error deleting products in database")

		if err.Error() == "record not found" {
			errResponse := requests.ErrorResponse("Product not found")
			return c.JSON(http.StatusNotFound, errResponse)
		}
		errResponse := requests.ErrorResponse("Internal Server Error")
		return c.JSON(http.StatusInternalServerError, errResponse)
	}

	productUpdated, err := h.Service.FindOne(id)
	if err != nil {
		errResponse := requests.ErrorResponse("Product not found")
		return c.JSON(http.StatusNotFound, errResponse)
	}

	log.Print("PUT :id request finished")
	successResponse := requests.SuccessResponse(*productUpdated)
	return c.JSON(http.StatusOK, successResponse)
}
