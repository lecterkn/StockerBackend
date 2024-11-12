package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type ItemStockRepository interface {
	Select(id uuid.UUID) (*entity.ItemStockEntity, error)
	Insert(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error)
}
