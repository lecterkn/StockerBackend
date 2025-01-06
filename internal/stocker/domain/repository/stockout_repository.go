package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type StockOutRepository interface {
	Index(storeId uuid.UUID) ([]entity.StockOutEntity, error)
	Create(entity *entity.StockOutEntity) (*entity.StockOutEntity, error)
}
