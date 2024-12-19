package model

import "time"

type StockOutModel struct {
	Id        []byte
	StoreId   []byte
	ItemId    []byte
	Place     *string
	Stocks    int
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (StockOutModel) TableName() string {
	return "stock_outs"
}

type StockOutQueryModel struct {
	Id            []byte
	StoreId       []byte
	ItemId        []byte
	Place         *string
	Stocks        int
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	JanCode       string
	ItemCreatedAt time.Time
	ItemUpdatedAt time.Time
}
