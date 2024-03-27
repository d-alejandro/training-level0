package memory_cache

import (
	"errors"
	"github.com/d-alejandro/training-level0/internal/models"
	"sync"
)

type Cache struct {
	sync.RWMutex
	models map[string]*models.Model
}

func NewCache(orderModels []*models.OrderModel) CacheInterface {
	cacheModels := make(map[string]*models.Model)

	for _, orderModel := range orderModels {
		cacheModels[orderModel.Model.OrderUID] = &orderModel.Model
	}

	return &Cache{models: cacheModels}
}

func (cache *Cache) GetModel(OrderUID string) (*models.Model, error) {
	cache.RLock()
	defer cache.RUnlock()

	model, found := cache.models[OrderUID]
	if !found {
		return &models.Model{}, errors.New("model not found")
	}

	return model, nil
}

func (cache *Cache) SetModel(OrderUID string, model *models.Model) error {
	cache.Lock()
	defer cache.Unlock()

	cache.models[OrderUID] = model

	return nil
}
