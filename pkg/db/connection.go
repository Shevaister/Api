package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if os.Getenv("DB_CONNECTION") == "" {
		os.Setenv("DB_CONNECTION", "mysql:mysql@tcp(127.0.0.1:3306)/testparse?charset=utf8&parseTime=True&loc=Local")
	}
	dsn := os.Getenv("DB_CONNECTION")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
