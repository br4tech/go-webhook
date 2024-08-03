package repositories

import (
	"errors"

	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type OrderItemRepository[T port.IModel] struct {
	orderItemAdapter port.IPostgreDatabase[T]
}

func NewOrderItemRepository[T port.IModel](orderItemAdapter port.IPostgreDatabase[T]) port.IOrderItemRepository {
	return &OrderItemRepository[T]{
		orderItemAdapter: orderItemAdapter,
	}
}

func (repo *OrderItemRepository[T]) Create(orderItem *domain.OrderItem) (*domain.OrderItem, error) {
	orderItemModel := new(model.OrderItem)
	orderItemModel.FromDomain(orderItem)

	orderItemEntity, ok := any(orderItemModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to Product")
	}

	orderItemId, err := repo.orderItemAdapter.Create(orderItemEntity)
	if err != nil {
		return nil, err
	}
	orderItem.Id = orderItemId

	return orderItem, nil
}
