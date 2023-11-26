package endpoint

import (
	"core"
	"github.com/gin-gonic/gin"
	"inputs"
	"net/http"
	"repositories"
	"transformers"
)

type Author struct {
	core.Endpoint
}

// Create godoc
// @Success 201
// @Accept json
// @Produce json
// @Router /api/authors [post]
func (e *Author) Create(ctx *gin.Context) {
	var input inputs.Author
	err := ctx.ShouldBind(&input)

	if err != nil {
		ctx.String(http.StatusBadRequest, "bad request : %v", err)
		return
	}

	author := new(transformers.Author).Transform(&input)
	new(repositories.Author).Create(&author)
	ctx.IndentedJSON(http.StatusCreated, author)
}

// Show godoc
// @Success 200
// @Router /api/authors/:id [get]
func (e *Author) Show(context *gin.Context) {
	id := context.Param("id")
	_, author := new(repositories.Author).Find(id)
	context.IndentedJSON(http.StatusOK, &author)
}
