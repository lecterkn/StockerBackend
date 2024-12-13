package entity

import (
	"time"

	"github.com/google/uuid"
)

type ItemStockEntity struct {
	Item      ItemEntity
	Price     *int
	Stock     int
	StockMin  *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewItemStockEntity(storeId uuid.UUID, name, janCode string, price *int, stock int, stockMin *int) (*ItemStockEntity, error) {
	item, err := NewItemEntity(storeId, name, janCode)
	if err != nil {
		return nil, err
	}
	return &ItemStockEntity{
		Item:      *item,
		Price:     price,
		Stock:     stock,
		StockMin:  stockMin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (entity *ItemStockEntity) Update(price *int, stock int, stockMin *int) {
	entity.Price = price
	entity.Stock = stock
	entity.StockMin = stockMin
	entity.UpdatedAt = time.Now()
}
