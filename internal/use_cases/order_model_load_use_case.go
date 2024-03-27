package use_cases

import "github.com/d-alejandro/training-level0/internal/models"

type OrderModelLoadUseCase struct {
	orderModelLoadRepository OrderModelLoadRepositoryInterface
}

func NewOrderModelLoadUseCase(orderModelLoadRepository OrderModelLoadRepositoryInterface) *OrderModelLoadUseCase {
	return &OrderModelLoadUseCase{orderModelLoadRepository}
}

func (orderModelLoadUseCase *OrderModelLoadUseCase) Execute() []*models.OrderModel {
	return orderModelLoadUseCase.orderModelLoadRepository.Make()
}
