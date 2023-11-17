package repositories

import (
	"core"
	"gorm.io/gorm"
	"models"
)

type Books struct {
	core.Repository
}

func (r *Books) FindAll() (*gorm.DB, []models.Book) {
	var books []models.Book
	result := r.Db().Find(&books)
	return result, books
}

func (r *Books) Find() (*gorm.DB, models.Book) {
	var book models.Book
	result := r.Db().First()
}
