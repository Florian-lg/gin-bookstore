package endpoint

import (
	"core"
	"db"
	"github.com/gin-gonic/gin"
	"inputs"
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
	id := ctx.Query("id")
	_, book := new(repositories.Books).Find(id)
	ctx.IndentedJSON(http.StatusOK, book)
}

func (e *Book) Create(ctx *gin.Context) {
	var book inputs.Book
	err := ctx.ShouldBind(&book)

	if err != nil {
		ctx.String(http.StatusBadRequest, "bad request : %v", err)
	}

	db.Db().Create(&book)
	ctx.IndentedJSON(http.StatusCreated, book)
}
