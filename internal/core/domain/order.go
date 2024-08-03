package domain

import "time"

type Order struct {
	Id             int       `validate:"required"`
	DateToDelivery time.Time `validate:"required"`
	Status         string    `validate:"required"`
	Items          []OrderItem
}

func NewOrder(id int, dateToDelivery time.Time, status string, items []OrderItem) *Order {
	return &Order{
		Id:             id,
		DateToDelivery: dateToDelivery,
		Status:         status,
		Items:          items,
	}
}
