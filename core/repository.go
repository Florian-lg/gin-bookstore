package core

import (
	"db"
	"gorm.io/gorm"
)

type Repository struct {
}

func (m *Repository) Db() *gorm.DB {
	return db.Db()
}
