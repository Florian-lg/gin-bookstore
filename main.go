package main

import (
	"db"
	"router"
)

func main() {
	db.Migrate()
	router.Routes()
}
