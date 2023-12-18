package models

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID `gorm:"column:id;not null; primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}