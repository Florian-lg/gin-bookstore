package repositories

import (
	"core"
	"gorm.io/gorm"
	"inputs"
	"models"
)

type Book struct {
	core.Repository
}

func (r *Book) Create(book *models.Book) {
	r.Db().Save(&book)
}

func (r *Book) Update(input *inputs.Book, model *models.Book) {
	r.Db().Save(&model)
}

func (r *Book) FindAll() (*gorm.DB, []models.Book) {
	var books []models.Book
	result := r.Db().Find(&books)
	return result, books
}

func (r *Book) Find(id string) (*gorm.DB, models.Book) {
	var book models.Book
	result := r.Db().First(&book, "id = ?", id)
	return result, book
}

func (r *Book) Destroy(id string) {
	var book models.Book
	r.Db().Delete(&book, "id = ?", id)
}
