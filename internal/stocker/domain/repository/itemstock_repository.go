package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type ItemStockRepository interface {
	Index(storeId uuid.UUID) ([]entity.ItemStockEntity, error)
	Select(storeId, id uuid.UUID) (*entity.ItemStockEntity, error)
	SelectByJancode(storeId uuid.UUID, jancode string) (*entity.ItemStockEntity, error)
	Insert(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error)
	Update(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error)
}
