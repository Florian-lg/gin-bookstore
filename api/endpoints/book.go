package endpoint

import (
	"core"
	"db"
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
	"repositories"
)

type Book struct {
	core.Endpoint
}

func (e *Book) Index(ctx *gin.Context) {
	_, books := new(repositories.Books).FindAll()
	ctx.IndentedJSON(http.StatusOK, books)
}

func (e *Book) Show(ctx *gin.Context) {
	book := new(repositories.Books).Find()
	ctx.IndentedJSON(http.StatusOK, book)
}

func (e *Book) Create(ctx *gin.Context) {
	book := models.Book{Title: "Test", Price: 10, Author: "toto"}
	db.Db().Create(&book)
	ctx.IndentedJSON(http.StatusCreated, book)
}
