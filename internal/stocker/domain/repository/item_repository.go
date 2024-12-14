package repository

import (
	"github.com/google/uuid"
	"h11/backend/internal/stocker/domain/entity"
)

type ItemRepository interface {
	Index(storeId uuid.UUID, jancode, name *string) ([]entity.ItemEntity, error)
	Select(storeId, id uuid.UUID) (*entity.ItemEntity, error)
	// TODO rename to Insert
	Create(entity *entity.ItemEntity) (*entity.ItemEntity, error)
	Update(entity *entity.ItemEntity) (*entity.ItemEntity, error)
}
