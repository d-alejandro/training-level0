package repositories

import (
	"github.com/d-alejandro/training-level0/internal/models"
	"gorm.io/gorm"
)

type OrderModelLoadRepository struct {
	gorm *gorm.DB
}

func NewOrderModelLoadRepository(gorm *gorm.DB) *OrderModelLoadRepository {
	return &OrderModelLoadRepository{gorm}
}

func (repository *OrderModelLoadRepository) Make() []*models.OrderModel {
	var orderModels []*models.OrderModel

	repository.gorm.Find(&orderModels)

	return orderModels
}
