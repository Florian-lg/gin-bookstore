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
		api.GET("/books/:id", books.Show)
		api.POST("/books", books.Create)
		api.DELETE("/books/:id", books.Destroy)
		api.PUT("/books/:id", books.Update)
	}

	err := r.Run()

	if err != nil {
		panic("Router error")
	}
}
