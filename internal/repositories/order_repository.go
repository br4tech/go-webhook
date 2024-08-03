package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type OrderRepository struct {
	postgres port.IPostgreDatabase[model.Order]
}

func NewOrderRepository(postgres port.IPostgreDatabase[model.Order]) port.IOrderRepository {
	return &OrderRepository{
		postgres: postgres,
	}
}

func (repo *OrderRepository) Create(order *domain.Order) (*domain.Order, error) {
	return nil, nil
}
