package dto

import "github.com/br4tech/go-webhook/internal/core/domain"

type SubscriptionDTO struct {
	Id           int `json:"id"`
	PaymentId    int `json:"payment_id"`
	SubscriberId int `json:"subscribe_id"`
}

func (dto *SubscriptionDTO) ToDomain() *domain.Subscription {
	return &domain.Subscription{
		Id:           dto.Id,
		PaymentId:    dto.PaymentId,
		SubscriberId: dto.SubscriberId,
	}
}

func (dto *SubscriptionDTO) FromDomain(domain *domain.Subscription) {
	dto.Id = domain.Id
	dto.PaymentId = domain.Id
	dto.SubscriberId = domain.SubscriberId
}
