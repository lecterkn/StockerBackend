package repository

import (
	"github.com/google/uuid"
	"h11/backend/internal/stocker/domain/entity"
)

type ItemRepository interface {
	SelectItems() ([]entity.ItemEntity, error)
	Select(id uuid.UUID) (*entity.ItemEntity, error)
	Create(entity *entity.ItemEntity) (*entity.ItemEntity, error)
	Update(entity *entity.ItemEntity) (*entity.ItemEntity, error)
}
