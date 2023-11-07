package main

import (
	"db"
	"router"
)

func main() {
	db.Setup()
	router.Routes()
}
