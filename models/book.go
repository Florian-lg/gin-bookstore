package models

import (
	"db"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func FindAllBooks() []Book {
	books := db.Setup().Find(&Book{})

	return books
}

func FindBookById(bookId string) *Book {
	for _, book := range books {
		if book.ID == bookId {
			return &book
		}
	}

	return nil
}
