package repository

import (
	"h11/backend/internal/stocker/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	Select(id uuid.UUID) (*entity.UserEntity, error)
	SelectByName(name string) (*entity.UserEntity, error)
    Insert(*entity.UserEntity) (*entity.UserEntity, error)
    Update(*entity.UserEntity) (*entity.UserEntity, error)
}