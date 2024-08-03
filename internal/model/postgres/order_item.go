package model

import "github.com/br4tech/go-webhook/internal/core/domain"

type OrderItem struct {
	Model
	Quantity  int     `gorm:"column:quantity"`
	ProductId int     `gorm:"column:product_id"`
	Price     float32 `gorm:"column:price"`
	SubPrice  float32 `gorm:"column:sub_price"`
}

func (model OrderItem) GetId() int {
	return model.Id
}

func (model *OrderItem) ToDomain() *domain.OrderItem {
	return &domain.OrderItem{
		Id:        model.Id,
		Quantity:  model.Quantity,
		ProductId: model.ProductId,
		Price:     model.Price,
		SubPrice:  model.SubPrice,
	}
}

func (model *OrderItem) FromDomain(domain *domain.OrderItem) {
	model.Id = domain.Id
	model.Quantity = domain.Quantity
	model.ProductId = domain.ProductId
	model.Price = domain.Price
	model.SubPrice = domain.SubPrice
}
