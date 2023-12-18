package endpoint

import (
	"core"
	"net/http"
	"inputs"
	"github.com/gin-gonic/gin"
)

type Security struct {
	core.Endpoint
}

// Register godoc
// @Success 201
// @Accept json
// @Produce json
// @Router /api/register [post]
func (e *Security) Register(ctx *gin.Context) {
	var input inputs.UserRegister
	
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "ok"})
}