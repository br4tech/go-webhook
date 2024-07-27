package model

import "github.com/br4tech/go-webhook/internal/core/domain"

type Subscription struct {
	Id           int `bson:"id"`
	PaymentId    int `bson:"payment_id"`
	SubscriberId int `bson:"susbscriber_id"`
}

func (model *Subscription) ToDomain() *domain.Subscription {
	return &domain.Subscription{
		Id:           model.Id,
		PaymentId:    model.PaymentId,
		SubscriberId: model.SubscriberId,
	}
}

func (model *Subscription) FromDomain(domain *domain.Subscription) {
	model.Id = domain.Id
	model.PaymentId = domain.PaymentId
	model.SubscriberId = domain.SubscriberId
}
