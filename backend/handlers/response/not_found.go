package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundResponse(context *gin.Context, message string) {
	context.JSON(http.StatusNotFound, gin.H{"error": message})
}
