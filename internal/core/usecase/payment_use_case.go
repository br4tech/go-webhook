package usecase

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
)

type PaymentUseCase struct {
	repo port.IPaymentRepository
}

func NewPaymentUsecase(repo port.IPaymentRepository) port.IPaymentUsecase {
	return &PaymentUseCase{
		repo: repo,
	}
}

func (usecase *PaymentUseCase) Create(payment *domain.Payment) (*domain.Payment, error) {
	pay, err := usecase.repo.Create(payment)
	if err != nil {
		return nil, err
	}

	return pay, nil
}
