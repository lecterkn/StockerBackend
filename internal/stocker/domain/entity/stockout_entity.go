package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockOutEntity struct {
	Id        uuid.UUID
	Price     int
	Stocks    int
	Item      ItemEntity
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStockOutEntity(item ItemEntity, price, stocks int) (*StockOutEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &StockOutEntity{
		Id:        id,
		Price:     price,
		Stocks:    stocks,
		Item:      item,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
