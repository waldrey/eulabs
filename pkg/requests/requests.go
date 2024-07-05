package requests

import "github.com/waldrey/eulabs/internal/entity"

type TypeErrorResponse struct {
	Data struct {
		Error string `json:"error"`
	} `json:"data"`
}

type TypeSuccessResponse struct {
	Data interface{} `json:"data"`
}

func ErrorResponse(message string) TypeErrorResponse {
	return TypeErrorResponse{
		Data: struct {
			Error string `json:"error"`
		}{
			Error: message,
		},
	}
}

func SuccessResponse(product entity.Product) TypeSuccessResponse {
	return TypeSuccessResponse{
		Data: product,
	}
}

func SuccessListResponse(products []entity.Product) TypeSuccessResponse {
	return TypeSuccessResponse{
		Data: products,
	}
}
