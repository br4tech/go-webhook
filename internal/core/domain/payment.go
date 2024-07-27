package domain

import "time"

type Payment struct {
	OrderId     string
	Description string
	Status      string
	UpdateAt    time.Time
}

func NewPayment(orderId, description, status string, updatedAt time.Time) *Payment {
	return &Payment{
		OrderId:     orderId,
		Description: description,
		Status:      status,
		UpdateAt:    updatedAt,
	}
}
