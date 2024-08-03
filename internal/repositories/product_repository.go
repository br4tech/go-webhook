package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/postgres"
)

type ProductRepository struct {
	productAdapter port.IPostgreDatabase[model.Product]
}

func NewProductRepository(productAdapter port.IPostgreDatabase[model.Product]) port.IProductRepository {
	return &ProductRepository{
		productAdapter: productAdapter,
	}
}

func (repo *ProductRepository) FindById(id int) (*domain.Product, error) {
	return nil, nil
}

func (repo *ProductRepository) Create(product *domain.Product) (*domain.Product, error) {
	productModel := new(model.Product)
	productModel.FromDomain(product)

	productId, err := repo.productAdapter.Create(*productModel)
	if err != nil {
		return nil, err
	}
	product.Id = productId

	return product, nil
}
