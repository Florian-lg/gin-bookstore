package models

import (
	"gorm.io/gorm"
)

type book struct {
	gorm.Model
	ID string `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
}

var books = []book{
	{ID: "0", Title: "Test 0", Price: 10},
	{ID: "1", Title: "Test 1", Price: 20},
	{ID: "1", Title: "Test 2", Price: 20},
}

func FindAllBooks() []book {
	return books
}

func FindBookById(bookId string) *book {
	for _, book := range books {
		if book.ID == bookId {
			return &book
		}
	}

	return nil
}