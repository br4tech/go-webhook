package dto

import "github.com/br4tech/go-webhook/internal/core/domain"

type SubscribeDTO struct {
	Id        int    `json:"id"`
	PaymentId int    `json:"payment_id"`
	Webhook   string `json:"weebhook"`
	Token     string `json:"token"`
}

func (dto *SubscribeDTO) ToDomain() *domain.Subscribe {
	return &domain.Subscribe{
		Id:        dto.Id,
		PaymentId: dto.PaymentId,
		Webhook:   dto.Webhook,
		Token:     dto.Token,
	}
}

func (dto *SubscribeDTO) FromDomain(domain *domain.Subscribe) {
	dto.Id = domain.Id
	dto.PaymentId = domain.Id
	dto.Webhook = domain.Webhook
	dto.Token = domain.Token
}
