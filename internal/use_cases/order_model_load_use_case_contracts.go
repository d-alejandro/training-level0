package use_cases

import "github.com/d-alejandro/training-level0/internal/models"

type OrderModelLoadRepositoryInterface interface {
	Make() []*models.OrderModel
}
