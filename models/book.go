package models

import "github.com/google/uuid"

type Book struct {
	Id       uuid.UUID `gorm:"column:id;not null;primaryKey" json:"id"`
	Title    string    `json:"title"`
	Price    float64   `json:"price"`
	AuthorID int
	Author   Author
	Timestamp
}
