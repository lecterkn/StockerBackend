package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockInEntity struct {
	Id        uuid.UUID
	Place     string
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Price     int
	Stocks    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStockInEntity(place string, storeId, itemId uuid.UUID, price, stocks int) (*StockInEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &StockInEntity{
		Id:        id,
		Place:     place,
		StoreId:   storeId,
		ItemId:    itemId,
		Stocks:    stocks,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
