package dto

import (
	"time"

	"github.com/br4tech/go-webhook/internal/core/domain"
)

type OrderDTO struct {
	DateToDelivery time.Time      `json:"date_delivery"`
	Status         string         `json:"status"`
	Items          []OrderItemDTO `json:"items"`
}

func (dto *OrderDTO) ToDomain() *domain.Order {
	orderDomain := &domain.Order{
		DateToDelivery: dto.DateToDelivery,
		Status:         dto.Status,
		Items:          make([]domain.OrderItem, len(dto.Items)),
	}

	for i, item := range dto.Items {
		orderDomain.Items[i] = *item.ToDomain()
	}

	return orderDomain
}

func (dto *OrderDTO) FromDomain(domain *domain.Order) {
	dto.DateToDelivery = domain.DateToDelivery
	dto.Status = domain.Status
	dto.Items = make([]OrderItemDTO, len(domain.Items))

	for i, item := range domain.Items {
		dto.Items[i].FromDomain(&item)
	}
}
