package handler

import (
	"net/http"

	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/dto"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase port.IProductUseCase
}

func NewProductHandler(productUseCase port.IProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (handler *ProductHandler) Create(c echo.Context) error {
	reqBody := new(dto.ProductDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := handler.productUseCase.Create(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create product")
	}

	return HandlerResponse(c, http.StatusOK, "Created product")
}
