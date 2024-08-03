package domain

type Product struct {
	Id   int    `validate:"required"`
	Code string `validate:"required"`
	Name string `validate:"required"`
}

func NewProduct(id int, code, name string) *Product {
	return &Product{
		Id:   id,
		Code: code,
		Name: name,
	}
}
