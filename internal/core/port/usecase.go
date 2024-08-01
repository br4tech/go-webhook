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
