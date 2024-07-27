package usecase

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
)

type SubscriptionUseCase struct {
	repo port.IRepository
}

func NewSubscriptionUseCase(repo port.IRepository) port.ISubscriptionUseCase {
	return &SubscriptionUseCase{
		repo: repo,
	}
}

func (usecase *SubscriptionUseCase) Create(subscriptrion *domain.Subscription) (*domain.Subscription, error) {
	sub, err := usecase.repo.Create(subscriptrion)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (usecase *SubscriptionUseCase) FindAll() ([]*domain.Subscription, error) {
	sub, err := usecase.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return sub, nil
}
