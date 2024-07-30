package port

import "github.com/br4tech/go-webhook/internal/core/domain"

type ISubscriptionRespository interface {
	FindAll() ([]domain.Subscription, error)
	Create(subscription *domain.Subscription) (*domain.Subscription, error)
}

type IPaymentRepository interface {
	Create(payment *domain.Payment) (*domain.Payment, error)
}
