package db

import (
	"log"
	"models"
)

func Migrate() {
	log.Printf("Starts migration")

	err := Db().AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.User{},
	)

	if err != nil {
		panic("Migration cannot be generated.")
	}
}
