package model

import "time"

type StockInModel struct {
	Id        []byte
	Place     string
	StoreId   []byte
	ItemId    []byte
	Stocks    int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StockInModel) TableName() string {
	return "stock_ins"
}
