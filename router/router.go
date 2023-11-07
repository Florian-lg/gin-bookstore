package router

import (
	"endpoint"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/books", endpoint.GetAllBooks)
	}

	err := r.Run()

	if err != nil {
		panic("Endpoint not found")
	}
}
