package domain

import "time"

type Order struct {
	Id             int       `validate:"required"`
	DateToDelivery time.Time `validate:"required"`
	Items          []OrderItem
}

func NewOrder(id int, dateToDelivery time.Time, items []OrderItem) *Order {
	return &Order{
		Id:             id,
		DateToDelivery: dateToDelivery,
		Items:          items,
	}
}
