package endpoint

import (
	"core"
	"inputs"
	"net/http"
	"repositories"
	"transformers"
	"github.com/gin-gonic/gin"
)

type Book struct {
	core.Endpoint
}

// Index godoc
// @Success 200
// @Produce json
// @Router /api/books [get]
func (e *Book) Index(ctx *gin.Context) {
	_, books := new(repositories.Book).FindAll()
	ctx.IndentedJSON(http.StatusOK, books)
}

// Show godoc
// @Success 200
// @Produce json
// @Router /api/books/:id [get]
func (e *Book) Show(ctx *gin.Context) {
	id := ctx.Param("id")
	result, book := new(repositories.Book).Find(id)

	if result == nil {
		ctx.IndentedJSON(http.StatusNotFound, "The requested book was not found.")
	}
	

	ctx.IndentedJSON(http.StatusOK, book)
}

// Create godoc
// @Success 201
// @Accept json
// @Produce json
// @Router /api/books [post]
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

// Update godoc
// @Success 201
// @Accept json
// @Produce json
// @Router /api/books [patch]
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

// Destroy godoc
// @Success 204
// @Accept json
// @Produce json
// @Router /api/books [delete]
func (e *Book) Destroy(ctx *gin.Context) {
	id := ctx.Param("id")
	new(repositories.Book).Destroy(id)
}
