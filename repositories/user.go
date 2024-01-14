package repositories

import (
	"core"
	"models"
)

type UserRepository struct {
	core.Repository
}

func (r *UserRepository) Create (user *models.User) {
	r.Db().Save(&user)
}