package model

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	OrderId  string             `bson:"order_id"`
	Status   string             `bason:"status"`
	UpdateAt primitive.DateTime `bson:"update_at"`
}

func (model *Payment) ToDomain() *domain.Payment {
	return &domain.Payment{
		OrderId:  model.OrderId,
		Status:   model.Status,
		UpdateAt: model.UpdateAt.Time(),
	}
}

func (model *Payment) FromDomain(domain *domain.Payment) {
	model.OrderId = domain.OrderId
	model.Status = domain.Status
	model.UpdateAt = primitive.NewDateTimeFromTime(domain.UpdateAt)
}
