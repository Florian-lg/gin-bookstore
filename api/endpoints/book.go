package endpoint

import (
	"core"
	"github.com/gin-gonic/gin"
	"inputs"
	"net/http"
	"repositories"
	"transformers"
)

type Book struct {
	core.Endpoint
}

func (e *Book) Index(ctx *gin.Context) {
	_, books := new(repositories.Book).FindAll()
	ctx.IndentedJSON(http.StatusOK, books)
}

func (e *Book) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	_, book := new(repositories.Book).Find(id)
	ctx.IndentedJSON(http.StatusOK, book)
}

func (e *Book) Create(ctx *gin.Context) {
	var input inputs.Book
	err := ctx.ShouldBind(&input)

	if err != nil {
		ctx.String(http.StatusBadRequest, "bad request : %v", err)
		return
	}

	book := new(transformers.Book).Transform(&input, nil)
	new(repositories.Book).Create(&book)
	ctx.IndentedJSON(http.StatusCreated, book)
}

func (e *Book) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	_, model := new(repositories.Book).Find(id)
	var input inputs.Book
	err := ctx.ShouldBind(&input)

	if err != nil {
		ctx.String(http.StatusBadRequest, "bad request : %v", err)
		return
	}

	book := new(transformers.Book).Transform(&input, &model)
	new(repositories.Book).Update(&input, &book)
	ctx.IndentedJSON(http.StatusCreated, book)
}

func (e *Book) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	new(repositories.Book).Destroy(id)
}
