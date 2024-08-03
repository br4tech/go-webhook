package handler

import "github.com/br4tech/go-webhook/internal/core/port"

type OrderHandler struct {
}

func NewOrderHandler() port.IOrderHandler {
	return &OrderHandler{}
}
