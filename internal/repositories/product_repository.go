package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type ProductRepository struct {
	postgres port.IPostgreDatabase[model.Product]
}

func NewProductRepository(postgres port.IPostgreDatabase[model.Product]) port.IProductRepository {
	return &ProductRepository{
		postgres: postgres,
	}
}

func (repo *ProductRepository) FindById(id int) (*domain.Product, error) {
	return nil, nil
}

func (repo *ProductRepository) Create(product *domain.Product) (*domain.Product, error) {
	productModel := new(model.Product)
	productModel.FromDomain(product)

	productId, err := repo.postgres.Create(*productModel)
	if err != nil {
		return nil, err
	}
	product.Id = productId

	return product, nil
}
