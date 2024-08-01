package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/model"
	"github.com/jinzhu/copier"
)

type SubscribeRespository struct {
	adapter port.IMongoDatabase[model.Subscribe]
}

func NewSubscribeRespository(adapter port.IMongoDatabase[model.Subscribe]) port.ISubscribeRespository {
	return &SubscribeRespository{
		adapter: adapter,
	}
}

func (repo *SubscribeRespository) FindAll() ([]domain.Subscribe, error) {
	subscribeDomain := []domain.Subscribe{}

	subscribes, err := repo.adapter.FindAll()
	if err != nil {
		return nil, err
	}

	err = copier.Copy(subscribeDomain, subscribes)

	return subscribeDomain, nil
}

func (repo *SubscribeRespository) Create(Subscribe *domain.Subscribe) (*domain.Subscribe, error) {
	SubscribeModel := new(model.Subscribe)
	SubscribeModel.FromDomain(Subscribe)

	err := repo.adapter.Create(*SubscribeModel)
	if err != nil {
		return nil, err
	}

	return Subscribe, nil
}
