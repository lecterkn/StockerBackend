package model

import "time"

type ItemStockModel struct {
	ItemId    []byte `gorm:"column:item_id"`
	StoreId   []byte `gorm:"column:store_id"`
	Price     *int
	Stock     int
	StockMin  int       `gorm:"column:stock_min"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ItemStockModel) TableName() string {
	return "item_stocks"
}

type ItemStockQueryModel struct {
	Name          string
	JanCode       string
	ItemCreatedAt time.Time
	ItemUpdatedAt time.Time
	ItemId        []byte `gorm:"column:item_id"`
	StoreId       []byte `gorm:"column:store_id"`
	Price         *int
	Stock         int
	StockMin      int       `gorm:"column:stock_min"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}
