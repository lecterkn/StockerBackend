package model

import "time"

type UserModel struct {
	Id        []byte
	Name      string
	Password  []byte
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (UserModel) TableName() string {
    return "users"
}