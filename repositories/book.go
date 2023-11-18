package repositories

import (
	"core"
	"gorm.io/gorm"
	"inputs"
	"models"
	"transformers"
)

type Books struct {
	core.Repository
}

func (r *Books) Create(input *inputs.Book) {
	book := transformers.Transform(input)
	r.Db().Create(&book)
}

func (r *Books) FindAll() (*gorm.DB, []models.Book) {
	var books []models.Book
	result := r.Db().Find(&books)
	return result, books
}

func (r *Books) Find(id string) (*gorm.DB, models.Book) {
	var book models.Book
	result := r.Db().First(&book, "id = ?", id)
	return result, book
}

func (r *Books) Destroy(id string) {
	var book models.Book
	r.Db().Delete(&book, "id = ?", id)
}
