package router

import (
	"endpoint"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

		authors := new(endpoint.Author)
		api.GET("/authors/:id", authors.Show)
		api.POST("/authors", authors.Create)
	}

	ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	err := r.Run()

	if err != nil {
		panic("Router error")
	}
}
