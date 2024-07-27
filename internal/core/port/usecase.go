package port

import "github.com/br4tech/go-webhook/internal/core/domain"

type ISubscriptionUseCase interface {
	Create(subscriptrion *domain.Subscription) (*domain.Subscription, error)
	FindAll() ([]*domain.Subscription, error)
}
