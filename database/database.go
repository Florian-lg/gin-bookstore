package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() {
	db, err := gorm.Open(sqlite.Open("bookstore.db", &gorm.Config{}))

	if err != nil {
		panic("Failed to connect to database")
	}
}
