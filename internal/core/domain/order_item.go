package domain

type OrderItem struct {
	Id        int     `validate:"required"`
	Quantity  int     `validate:"required"`
	ProductId int     `validate:"required"`
	Price     float32 `validate:"required"`
	SubPrice  float32 `validate:"required"`
}

func NewOrderItem(id, quantity, productId int, price, subprice float32) *OrderItem {
	return &OrderItem{
		Id:        id,
		Quantity:  quantity,
		ProductId: productId,
		Price:     price,
		SubPrice:  subprice,
	}
}
