package repository

import "h11/backend/internal/stocker/domain/entity"

type ItemRepository interface {
	SelectItems() ([]entity.ItemEntity, error)
	Create(entity *entity.ItemEntity) (*entity.ItemEntity, error)
}
