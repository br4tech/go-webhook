package repositories

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	model "github.com/br4tech/go-webhook/internal/model/mongo"
	"github.com/br4tech/go-webhook/internal/utils/filters"
	"github.com/jinzhu/copier"
)

type PaymentRepository struct {
	adapter port.IMongoDatabase[model.Payment]
}

func NewPaymentRepository(adapter port.IMongoDatabase[model.Payment]) port.IPaymentRepository {
	return &PaymentRepository{
		adapter: adapter,
	}
}

func (repo *PaymentRepository) FindByOrdeId(orderId string) ([]domain.Payment, error) {
	paymentDomain := []domain.Payment{}

	filter := filters.EqualFilter{
		Field: orderId,
		Value: orderId,
	}

	payments, err := repo.adapter.FindBy(filter)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(paymentDomain, payments)

	return paymentDomain, nil
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
