package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type OrderItemRepository struct {
	postgres port.IPostgreDatabase[model.OrderItem]
}

func NewOrderItemRepository(postgres port.IPostgreDatabase[model.OrderItem]) port.IOrderItemRepository {
	return &OrderItemRepository{
		postgres: postgres,
	}
}

func (repo *OrderItemRepository) Create(orderItem *domain.OrderItem) (*domain.OrderItem, error) {
	return nil, nil
}
