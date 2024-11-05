package model

import (
	"time"
)

type ItemModel struct {
	Id        []byte `gorm:"type:uuid"`
	Name      string
	JanCode   string    `gorm:"column:jan_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (ItemModel) TableName() string {
	return "items"
}
