package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type StoreRepository interface {
	Index(userId uuid.UUID) ([]entity.StoreEntity, error)
	Select(id uuid.UUID) (*entity.StoreEntity, error)
	Insert(entity *entity.StoreEntity) (*entity.StoreEntity, error)
	Update(entity *entity.StoreEntity) (*entity.StoreEntity, error)
}
