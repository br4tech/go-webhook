package dto

import "github.com/br4tech/go-webhook/internal/core/domain"

type SubscribeDTO struct {
	Id           int `json:"id"`
	PaymentId    int `json:"payment_id"`
	SubscriberId int `json:"subscribe_id"`
}

func (dto *SubscribeDTO) ToDomain() *domain.Subscribe {
	return &domain.Subscribe{
		Id:           dto.Id,
		PaymentId:    dto.PaymentId,
		SubscriberId: dto.SubscriberId,
	}
}

func (dto *SubscribeDTO) FromDomain(domain *domain.Subscribe) {
	dto.Id = domain.Id
	dto.PaymentId = domain.Id
	dto.SubscriberId = domain.SubscriberId
}
