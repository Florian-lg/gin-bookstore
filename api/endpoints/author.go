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
