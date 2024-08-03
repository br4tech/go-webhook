package usecase

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
)

type ProductUseCase struct {
	productRepo port.IProductRepository
}

func NewProductUseCase(productRepo port.IProductRepository) port.IProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (usecase *ProductUseCase) FindById(id int) (*domain.Product, error) {
	product, err := usecase.productRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (usecase *ProductUseCase) Create(product *domain.Product) (*domain.Product, error) {
	product, err := usecase.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
