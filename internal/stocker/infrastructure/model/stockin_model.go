package model

import "time"

type StockInModel struct {
	Id        []byte
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

type StockInQueryModel struct {
	Id            []byte
	StoreId       []byte
	ItemId        []byte
	Stocks        int
	Price         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	JanCode       string
	ItemCreatedAt time.Time
	ItemUpdatedAt time.Time
}
