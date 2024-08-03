package repositories

import (
	"errors"

	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type OrderRepository[T port.IModel] struct {
	orderAdapter port.IPostgreDatabase[T]
}

func NewOrderRepository[T port.IModel](orderAdapter port.IPostgreDatabase[T]) port.IOrderRepository {
	return &OrderRepository[T]{
		orderAdapter: orderAdapter,
	}
}

func (repo *OrderRepository[T]) Create(order *domain.Order) (*domain.Order, error) {
	orderModel := new(model.Order)
	orderModel.FromDomain(order)

	orderEntity, ok := any(orderModel).(T)
	if !ok {
		return nil, errors.New("entity type invalid to Order")
	}

	orderId, err := repo.orderAdapter.Create(orderEntity)
	if err != nil {
		return nil, err
	}
	order.Id = orderId

	return order, nil
}
