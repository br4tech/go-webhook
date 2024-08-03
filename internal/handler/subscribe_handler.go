package handler

import (
	"net/http"

	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/dto"
	"github.com/labstack/echo/v4"
)

type SubscribeHandler struct {
	usecase port.ISubscribeUseCase
}

func NewSubscribeHandler(usecase port.ISubscribeUseCase) port.ISubscribeHandler {
	return &SubscribeHandler{
		usecase: usecase,
	}
}

func (handler *SubscribeHandler) Create(c echo.Context) error {
	reqBody := new(dto.SubscribeDTO)

	if err := c.Bind(reqBody); err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Bad request")
	}

	_, err := handler.usecase.Create(reqBody.ToDomain())
	if err != nil {
		return HandlerResponse(c, http.StatusInternalServerError, "Failed to create subscriptiopn")
	}

	return HandlerResponse(c, http.StatusOK, "Created Subscribe")
}