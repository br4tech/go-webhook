package model

import (
	"time"

	"github.com/br4tech/go-webhook/internal/core/domain"
)

type Order struct {
	Model
	DateToDelivery time.Time   `gorm:"colunm:date_delivery"`
	Status         string      `gorm:"colunm:status"`
	Items          []OrderItem `gorm:"foreignKey:OrderId"`
}

func (model Order) GetId() int {
	return model.Id
}

func (model *Order) ToDomain() *domain.Order {
	orderDomain := &domain.Order{
		DateToDelivery: model.DateToDelivery,
		Status:         model.Status,
		Items:          make([]domain.OrderItem, len(model.Items)),
	}

	for i, item := range model.Items {
		orderDomain.Items[i] = *item.ToDomain()
	}

	return orderDomain

}

func (model *Order) FromDomain(domain *domain.Order) {
	model.Id = domain.Id
	model.DateToDelivery = domain.DateToDelivery
	model.Status = domain.Status
	model.Items = make([]OrderItem, len(domain.Items))

	for i, item := range domain.Items {
		model.Items[i].FromDomain(&item)
	}
}
