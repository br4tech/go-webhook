package model

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
)

type Product struct {
	Model
	Code string `gorm:"unique;not null;column:code"`
	Name string `gorm:"unique;not null;column:name"`
}

func (model Product) GetId() int {
	return model.Id
}

func (model *Product) ToDomain() *domain.Product {
	return &domain.Product{
		Id:   model.Id,
		Code: model.Code,
		Name: model.Name,
	}
}

func (model *Product) FromDomain(domain *domain.Product) {
	model.Id = domain.Id
	model.Code = domain.Code
	model.Name = domain.Name
}
