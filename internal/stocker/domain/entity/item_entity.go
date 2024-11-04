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

func (entity *ItemEntity) Update(name, janCode string) {
    entity.Name = name
    entity.JanCode = janCode
    entity.UpdatedAt = time.Now()
}
