package router

import (
	"endpoint"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	api := r.Group("/api")
	{
		books := new(endpoint.Book)
		api.GET("/books", books.Index)
	}

	err := r.Run()

	if err != nil {
		panic("Endpoint not found")
	}
}
