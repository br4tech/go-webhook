package model

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Id   int    `gorm:"unique;not null;column:id"`
	Code string `gorm:"unique;not null;column:code"`
	Name string `gorm:"unique;not null;column:name"`
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
