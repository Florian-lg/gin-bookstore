package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
