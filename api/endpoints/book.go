package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"repositories"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repositories.Book{}.Find)
}
