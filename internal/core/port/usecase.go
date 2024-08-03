package port

import "github.com/br4tech/go-webhook/internal/core/domain"

type ISubscribeUseCase interface {
	Create(subscriptrion *domain.Subscribe) (*domain.Subscribe, error)
	FindAll() ([]domain.Subscribe, error)
}

type IPaymentUsecase interface {
	FindByOrderId(orderId string) ([]domain.Payment, error)
	Create(payment *domain.Payment) (*domain.Payment, error)
}

type IProductUseCase interface {
	FindById(id int) (*domain.Product, error)
	Create(product *domain.Product) (*domain.Product, error)
}

type IOrderUseCase interface {
	Create(order *domain.Order) (*domain.Order, error)
}
