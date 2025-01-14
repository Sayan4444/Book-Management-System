package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := os.Getenv("DB")
	if dsn == "" {
		panic("DB environment variable not set")
	}
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
	log.Printf("DB Connected")
}

func GetDB() *gorm.DB {
	return db
}
