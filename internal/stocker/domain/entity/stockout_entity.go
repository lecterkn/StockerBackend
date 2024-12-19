package entity

import (
	"time"

	"github.com/google/uuid"
)

type StockOutEntity struct {
	Id        uuid.UUID
	Place     *string
	Price     int
	Stocks    int
	Item      ItemEntity
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewStockOutEntity(item ItemEntity, place *string, price, stocks int) (*StockOutEntity, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &StockOutEntity{
		Id:        id,
		Place:     place,
		Price:     price,
		Stocks:    stocks,
		Item:      item,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
