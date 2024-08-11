package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Define the connection string
	dsn := "root:root@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
