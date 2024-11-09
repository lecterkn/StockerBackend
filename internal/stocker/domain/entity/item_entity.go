package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemEntity struct {
	Id        uuid.UUID
	Name      string
	JanCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewItemEntity /* エンティティのインスタンスを作成
func NewItemEntity(name, janCode string) (*ItemEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &ItemEntity{
		Id:        id,
		Name:      name,
		JanCode:   janCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Update /* エンティティを更新する
func (entity *ItemEntity) Update(name, janCode string) {
	entity.Name = name
	entity.JanCode = janCode
	entity.UpdatedAt = time.Now()
}
