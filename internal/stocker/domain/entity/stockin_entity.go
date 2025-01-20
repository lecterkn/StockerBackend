package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockInEntity struct {
	Id        uuid.UUID
	Price     int
	Stocks    int
	Item      ItemEntity
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStockInEntity(item ItemEntity, price, stocks int) (*StockInEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &StockInEntity{
		Id:        id,
		Item:      item,
		Stocks:    stocks,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
