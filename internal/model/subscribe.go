package model

import "github.com/br4tech/go-webhook/internal/core/domain"

type Subscribe struct {
	Id           int `bson:"id"`
	PaymentId    int `bson:"payment_id"`
	SubscriberId int `bson:"susbscriber_id"`
}

func (model *Subscribe) ToDomain() *domain.Subscribe {
	return &domain.Subscribe{
		Id:           model.Id,
		PaymentId:    model.PaymentId,
		SubscriberId: model.SubscriberId,
	}
}

func (model *Subscribe) FromDomain(domain *domain.Subscribe) {
	model.Id = domain.Id
	model.PaymentId = domain.PaymentId
	model.SubscriberId = domain.SubscriberId
}
