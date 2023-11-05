package main

import (
	"github.com/gin-gonic/gin"

	"endpoint"
)

func main() {
	r := gin.Default()

	r.GET("/books", endpoint.GetAllBooks)
	r.GET("/books/:bookId", endpoint.GetBook)

	r.Run()
}
