package model

import "time"

type ItemStockModel struct {
	ItemId    []byte `gorm:"column:item_id"`
	StoreId []byte `gorm:"column:store_id"`
	Place     string
	Stock     int
	StockMin  int       `gorm:"column:stock_min"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ItemStockModel) TableName() string {
	return "item_stocks"
}
