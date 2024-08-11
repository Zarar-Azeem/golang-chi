package models

import (

	database "github.com/Zarar-Azeem/golang-chi/pkg/db"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model `json:"-"` // Exclude gorm.Model fields from JSON encoding

	Name        string `json:"name" gorm:"type:varchar(255); not null"`
	Author      string `json:"author" gorm:"type:varchar(255); not null"`
	Publication string `json:"publication" gorm:"type:varchar(255)"`
}

func Init() {
	database.Connect()
	db = database.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) GetAllBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	return books, result.Error
}

func (b *Book) AddBook(book Book) error  {
	result := db.Create(&book)
	return result.Error
}
