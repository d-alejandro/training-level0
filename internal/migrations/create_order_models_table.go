package migrations

import (
	"github.com/d-alejandro/training-level0/internal/models"
	"gorm.io/gorm"
)

type OrderModelsTableMigration struct {
	gorm *gorm.DB
}

func NewOrderModelsTableMigration(gorm *gorm.DB) *OrderModelsTableMigration {
	return &OrderModelsTableMigration{gorm}
}

func (migration *OrderModelsTableMigration) Migrate() error {
	return migration.gorm.AutoMigrate(&models.OrderModel{})
}
