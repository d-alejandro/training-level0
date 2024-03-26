package use_cases

import (
	"github.com/d-alejandro/training-level0/internal/models"
)

type OrderModelStoreRepositoryInterface interface {
	Make(*models.Model) error
}
