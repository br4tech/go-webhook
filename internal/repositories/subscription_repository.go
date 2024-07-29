package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/model"
	"github.com/jinzhu/copier"
)

type SubscriptionRespository struct {
	adapter port.IMongoDatabase[model.Subscription]
}

func NewSubscriptionRespository(adapter port.IMongoDatabase[model.Subscription]) port.ISubscriptionRespository {
	return &SubscriptionRespository{
		adapter: adapter,
	}
}

func (repo *SubscriptionRespository) FindAll() ([]domain.Subscription, error) {
	subscriptionDomain := []domain.Subscription{}

	subscriptions, err := repo.adapter.FindAll()
	if err != nil {
		return nil, err
	}

	err = copier.Copy(subscriptionDomain, subscriptions)

	return subscriptionDomain, nil
}

func (repo *SubscriptionRespository) Create(subscription *domain.Subscription) (*domain.Subscription, error) {
	subscriptionModel := new(model.Subscription)
	subscriptionModel.FromDomain(subscription)

	err := repo.adapter.Create(*subscriptionModel)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
