package repositories

import (
	"github.com/d-alejandro/training-level0/internal/models"
	"gorm.io/gorm"
)

type OrderModelStoreRepository struct {
	gorm *gorm.DB
}

func NewOrderModelStoreRepository(gorm *gorm.DB) *OrderModelStoreRepository {
	return &OrderModelStoreRepository{gorm}
}

func (orderModelStoreRepository *OrderModelStoreRepository) Make(model *models.Model) error {
	orderModel := &models.OrderModel{Model: *model}
	result := orderModelStoreRepository.gorm.Create(orderModel)
	return result.Error
}
