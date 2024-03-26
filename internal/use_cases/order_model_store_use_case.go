package use_cases

import (
	"github.com/d-alejandro/training-level0/internal/models"
)

type OrderModelStoreUseCase struct {
	orderModelStoreRepository OrderModelStoreRepositoryInterface
}

func NewOrderModelStoreUseCase(orderModelStoreRepository OrderModelStoreRepositoryInterface) *OrderModelStoreUseCase {
	return &OrderModelStoreUseCase{orderModelStoreRepository}
}

func (orderModelStoreUseCase *OrderModelStoreUseCase) Execute(model *models.Model) error {
	return orderModelStoreUseCase.orderModelStoreRepository.Make(model)
}
