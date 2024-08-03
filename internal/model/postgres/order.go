package model

import (
	"time"

	"github.com/br4tech/go-webhook/internal/core/domain"
)

type Order struct {
	Id             int         `gorm:"primaryKey"`
	DateToDelivery time.Time   `gorm:"colunm:date_delivery"`
	Items          []OrderItem `gorm:"foreignKey:OrderId"`
}

func (model *Order) ToDomain() *domain.Order {
	return &domain.Order{
		Id:             model.Id,
		DateToDelivery: model.DateToDelivery,
		// Items:          model.Items,
	}
}

func (model *Order) FromDomain(domain *domain.Order) {
	model.Id = domain.Id
	model.DateToDelivery = domain.DateToDelivery
	// model.Items = domain.Items
}
