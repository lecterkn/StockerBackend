package model

import "time"

type StoreModel struct {
	Id        []byte
	UserId    []byte `gorm:"column:user_id"`
	Name      string
	CreatedAt time.Time
    UpdatedAt time.Time
}

func (StoreModel) TableName() string {
    return "stores"
}