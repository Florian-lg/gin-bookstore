package endpoint

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"models"
)

func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.FindAllBooks())
}

func GetBook(c *gin.Context) {
	bookId := c.Param("bookId")
	books := models.FindBookById(bookId)

	if books == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, models.FindBookById(bookId))
}