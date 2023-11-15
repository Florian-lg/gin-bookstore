package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Author string  `json:"author"`
}
