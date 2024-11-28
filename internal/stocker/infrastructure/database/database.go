package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetMySQLConnection() *gorm.DB {
    dsn, ok := os.LookupEnv("DATABASE_URL")
    if !ok {
        panic("\"DATABASE_URL\" is not set!")
    }
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
    return db
}
