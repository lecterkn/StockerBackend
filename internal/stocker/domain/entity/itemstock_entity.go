package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemStockEntity struct {
	ItemId    uuid.UUID
	StoreId   uuid.UUID
	Place     string
	Price     int
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemStockEntity(itemId, storeId uuid.UUID, place string, price, stock, stockMin int) (*ItemStockEntity, error) {
	return &ItemStockEntity{
		ItemId:    itemId,
		StoreId:   storeId,
		Place:     place,
		Price:     price,
		Stock:     stock,
		StockMin:  stockMin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (entity *ItemStockEntity) Update(place string, price, stock, stockMin int) {
	entity.Place = place
	entity.Price = price
	entity.Stock = stock
	entity.StockMin = stockMin
	entity.UpdatedAt = time.Now()
}
