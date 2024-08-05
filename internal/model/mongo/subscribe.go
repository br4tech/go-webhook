package model

import "github.com/br4tech/go-webhook/internal/core/domain"

type Subscribe struct {
	Id        int    `bson:"id"`
	PaymentId int    `bson:"payment_id"`
	Webhook   string `bson:"webhook"`
	Token     string `bson:"token"`
}

func (model *Subscribe) ToDomain() *domain.Subscribe {
	return &domain.Subscribe{
		Id:        model.Id,
		PaymentId: model.PaymentId,
		Webhook:   model.Webhook,
		Token:     model.Token,
	}
}

func (model *Subscribe) FromDomain(domain *domain.Subscribe) {
	model.Id = domain.Id
	model.PaymentId = domain.PaymentId
	model.Webhook = domain.Webhook
	model.Token = domain.Token
}
