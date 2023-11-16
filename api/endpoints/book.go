package endpoint

import (
	"core"
	"github.com/gin-gonic/gin"
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
