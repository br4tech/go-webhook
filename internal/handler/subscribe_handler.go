package handler

import (
	"net/http"

	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/dto"
	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	usecase port.ISubscriptionUseCase
}

func NewSubscriptionHandler(usecase port.ISubscriptionUseCase) port.ISubscriptionHandler {
	return &SubscriptionHandler{
		usecase: usecase,
	}
}

func (handler *SubscriptionHandler) Create(c echo.Context) error {
	reqBody := new(dto.SubscriptionDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := handler.usecase.Create(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create subscriptiopn")
	}

	return HandlerResponse(c, http.StatusOK, "Created subscription")
}
