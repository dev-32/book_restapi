package config

import (
	"github.com/jinzhu/gorm"
)
import _ "github.com/go-sql-driver/mysql"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "golden:password@/fresh_books?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
