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
	result, list := new(repositories.Books).FindAll()

	if result != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
	}

	ctx.IndentedJSON(http.StatusOK, list)

}
