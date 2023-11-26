package repositories

import (
	"core"
	"gorm.io/gorm"
	"models"
)

type Author struct {
	core.Repository
}

func (r *Author) Create(author *models.Author) {
	r.Db().Save(&author)
}

func (r *Author) Find(id string) (*gorm.DB, models.Author) {
	var author models.Author
	result := r.Db().First(&author, "id = ?", id)
	return result, author
}
