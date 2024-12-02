package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemStockEntity struct {
	ItemId    uuid.UUID
	StoreId uuid.UUID
	Place     string
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemStockEntity(itemId, storeId uuid.UUID, place string, stock, stockMin int) (*ItemStockEntity, error) {
	return &ItemStockEntity{
		ItemId:    itemId,
		StoreId: storeId,
		Place:     place,
		Stock:     stock,
		StockMin:  stockMin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (entity *ItemStockEntity) Update(place string, stock, stockMin int) {
	entity.Place = place
	entity.Stock = stock
	entity.StockMin = stockMin
	entity.UpdatedAt = time.Now()
}
