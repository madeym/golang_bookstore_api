package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "dev:Atomic100%@tcp(127.0.0.1:3306)/belajar-api?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB {
	return DB
}
