package handler

import (
	"net/http"

	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/dto"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	usecase port.IPaymentUsecase
}

func NewPaymnentHandler(usecase port.IPaymentUsecase) port.IPaymentHandler {
	return &PaymentHandler{
		usecase: usecase,
	}
}

func (handler *PaymentHandler) Create(c echo.Context) error {
	reqBody := new(dto.PaymentDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := handler.usecase.Create(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create subscriptiopn")
	}

	return HandlerResponse(c, http.StatusOK, "Created payment")
}
