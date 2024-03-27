package memory_cache

import "github.com/d-alejandro/training-level0/internal/models"

type CacheInterface interface {
	GetModel(OrderUID string) (*models.Model, error)
	SetModel(OrderUID string, model *models.Model) error
}
