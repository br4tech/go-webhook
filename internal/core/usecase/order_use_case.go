package usecase

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
)

type OrderUsecase struct {
	orderRepository     port.IOrderRepository
	orderItemReposiotry port.IOrderItemRepository
}

func NewOrderUseCase(orderRepository port.IOrderRepository,
	orderItemReposiotry port.IOrderItemRepository) *OrderUsecase {
	return &OrderUsecase{
		orderRepository:     orderRepository,
		orderItemReposiotry: orderItemReposiotry,
	}
}

func (usecase *OrderUsecase) Create(order *domain.Order) (*domain.Order, error) {
	return nil, nil
}
