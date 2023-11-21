package db

import (
	"log"
	"models"
)

func Migrate() {
	log.Printf("Starts migration")
	Db().AutoMigrate(
		&models.Book{},
		&models.Author{},
	)
}
