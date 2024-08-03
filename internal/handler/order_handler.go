package handler

import (
	"net/http"

	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/dto"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUseCase port.IOrderUseCase
}

func NewOrderHandler(orderUseCase port.IOrderUseCase) port.IOrderHandler {
	return &OrderHandler{
		orderUseCase: orderUseCase,
	}
}

func (handler *OrderHandler) Create(c echo.Context) error {
	reqBody := new(dto.OrderDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := handler.orderUseCase.Create(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create order")
	}

	return HandlerResponse(c, http.StatusOK, "Created order")
}
