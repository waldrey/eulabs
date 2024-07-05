package tools

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/waldrey/eulabs/pkg/requests"
)

func ValidateRequest(c echo.Context) (int, error) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errResponse := requests.ErrorResponse("ID must be an integer")
		return 0, c.JSON(http.StatusBadRequest, errResponse)
	}

	if id <= 0 {
		errResponse := requests.ErrorResponse("ID must be a positive integer")
		return 0, c.JSON(http.StatusBadRequest, errResponse)
	}

	return id, nil
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("the field '%s' is %s", strings.ToLower(err.Field()), err.Tag()))
	}

	return errors
}

func SafeDereferenceString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func SafeDereferenceFloat64(ptr *float64) float64 {
	if ptr == nil {
		return 0.0
	}
	return *ptr
}
