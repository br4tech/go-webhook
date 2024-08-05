package usecase

import (
	"github.com/br4tech/go-webhook/internal/core/domain"
	"github.com/br4tech/go-webhook/internal/core/port"
	"github.com/br4tech/go-webhook/internal/utils/validator"
)

type SubscribeUseCase struct {
	repo port.ISubscribeRespository
}

func NewSubscribeUseCase(repo port.ISubscribeRespository) port.ISubscribeUseCase {
	return &SubscribeUseCase{
		repo: repo,
	}
}

func (usecase *SubscribeUseCase) Create(subscriptrion *domain.Subscribe) (*domain.Subscribe, error) {

	if err := validator.ValidateStruct(subscriptrion); err != nil {
		return nil, err
	}

	sub, err := usecase.repo.Create(subscriptrion)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (usecase *SubscribeUseCase) FindAll() ([]domain.Subscribe, error) {
	sub, err := usecase.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return sub, nil
}
