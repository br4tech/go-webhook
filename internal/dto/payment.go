package dto

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
)

type PaymentDTO struct {
	OrderId     string `json:"order_id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (dto *PaymentDTO) ToDomain() *domain.Payment {
	return &domain.Payment{
		OrderId:     dto.OrderId,
		Description: dto.Description,
		Status:      dto.Status,
	}
}

func (dto *PaymentDTO) FromDomain(domain *domain.Payment) {
	dto.OrderId = domain.OrderId
	dto.Description = domain.Description
	dto.Status = domain.Status
}
