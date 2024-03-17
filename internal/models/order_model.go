package models

import (
	"github.com/d-alejandro/training-level0/internal/models/types"
	"time"
)

type OrderModel struct {
	ID        uint             `gorm:"primaryKey"`
	Model     types.ModelJSONB `gorm:"type:jsonb;not null;index:idx_order_models_model_order_uid,unique,expression:(model->>'order_uid')"`
	CreatedAt *time.Time       `gorm:"type:timestamp(0)"`
	UpdatedAt *time.Time       `gorm:"type:timestamp(0)"`
}
