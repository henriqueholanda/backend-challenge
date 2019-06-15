package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestResponse(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, gin.H{"error": message})
}
