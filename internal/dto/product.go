package dto

import "github.com/br4tech/go-webhook/internal/core/domain"

type ProductDTO struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (dto *ProductDTO) ToDomain() *domain.Product {
	return &domain.Product{
		Id:   dto.Id,
		Code: dto.Code,
		Name: dto.Name,
	}
}

func (dto *ProductDTO) FromDomain(domain domain.Product) {
	dto.Id = domain.Id
	dto.Code = domain.Code
	dto.Name = domain.Name
}
