package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/model"
)

type PaymentRepository struct {
	adapter port.IMongoDatabase[model.Payment]
}

func NewPaymentRepository(adapter port.IMongoDatabase[model.Payment]) port.IPaymentRepository {
	return &PaymentRepository{
		adapter: adapter,
	}
}

func (repo *PaymentRepository) Create(payment *domain.Payment) (*domain.Payment, error) {
	paymentModel := new(model.Payment)
	paymentModel.FromDomain(payment)

	err := repo.adapter.Create(*paymentModel)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
