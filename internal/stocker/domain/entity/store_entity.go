package entity

import (
	"time"

	"github.com/google/uuid"
)

type StoreEntity struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStoreEntity(userId uuid.UUID, name string) (*StoreEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &StoreEntity{
		Id:        id,
		UserId:    userId,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (entity *StoreEntity) Update(name string) {
	entity.Name = name
	entity.UpdatedAt = time.Now()
}
