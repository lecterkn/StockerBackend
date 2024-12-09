package entity

import (
	"h11/backend/internal/stocker/common"
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	Id        uuid.UUID
	Name      string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserEntity(name, password string) (*UserEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	pass, err := common.GetHashedPassword(password)
	if err != nil {
		return nil, err
	}
	return &UserEntity{
		Id:        id,
		Name:      name,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
