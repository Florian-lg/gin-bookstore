package models

import "github.com/google/uuid"

type Author struct {
	Id        uuid.UUID `gorm:"column:id;not null;primaryKey" json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Books     []Book    `json:"books"`
	Timestamp
}
