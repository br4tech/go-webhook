package port

import "github.com/br4tech/go-webhook/internal/core/domain"

type ISubscribeRespository interface {
	FindAll() ([]domain.Subscribe, error)
	Create(Subscribe *domain.Subscribe) (*domain.Subscribe, error)
}

type IPaymentRepository interface {
	FindByOrdeId(orderId string) ([]domain.Payment, error)
	Create(payment *domain.Payment) (*domain.Payment, error)
}

type IProductRepository interface {
	FindById(id int) (*domain.Product, error)
	Create(product *domain.Product) (*domain.Product, error)
}

type IOrderRepository interface {
	Create(order *domain.Order) (*domain.Order, error)
}

type IOrderItemRepository interface {
	Create(orderItem *domain.OrderItem) (*domain.OrderItem, error)
}
