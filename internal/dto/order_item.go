package dto

import "github.com/br4tech/go-webhook/internal/core/domain"

type OrderItemDTO struct {
	Quantity  int     `json:"quantity"`
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
	SubPrice  float32 `json:"sub_price"`
}

func (dto *OrderItemDTO) ToDomain() *domain.OrderItem {
	return &domain.OrderItem{
		Quantity:  dto.Quantity,
		ProductId: dto.ProductId,
		Price:     dto.Price,
		SubPrice:  dto.SubPrice,
	}
}

func (dto *OrderItemDTO) FromDomain(domain *domain.OrderItem) {
	dto.Quantity = domain.Quantity
	dto.ProductId = domain.ProductId
	dto.Price = domain.Price
	dto.SubPrice = domain.SubPrice
}
