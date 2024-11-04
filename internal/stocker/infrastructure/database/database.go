package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB = nil

func GetMySQLConnection() *gorm.DB {
	if connector == nil {
		dsn := os.Getenv("DATABASE_URL")
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil
		}
		connector = db
	}
	return connector
}
