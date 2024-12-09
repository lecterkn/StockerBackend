package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type StockInRepository interface {
	Create(entity *entity.StockInEntity) (*entity.StockInEntity, error)
	Index(storeId uuid.UUID) ([]entity.StockInEntity, error)
}
