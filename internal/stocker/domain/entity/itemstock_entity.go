package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemStockEntity struct {
	ItemId    uuid.UUID
	Place     string
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemStockEntity(itemId uuid.UUID, place string, stock, stockMin int) (*ItemStockEntity, error) {
	return &ItemStockEntity{
		ItemId:    itemId,
		Place:     place,
		Stock:     stock,
		StockMin:  stockMin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (entity ItemStockEntity) Update(place string, stock, stockMin int) *ItemStockEntity {
	return &ItemStockEntity{
		ItemId:    entity.ItemId,
		Place:     place,
		Stock:     stock,
		StockMin:  stockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: time.Now(),
	}
}
