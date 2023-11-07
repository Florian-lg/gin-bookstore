package router

import(
	"github.com/gin-gonic/gin"

	"endpoint"
)

func Routes() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/books", endpoint.GetAllBooks)
		api.GET("/books/:id", endpoint.GetBook)
	}

	r.Run()
}