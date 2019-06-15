package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatedResponse(context *gin.Context, content gin.H) {
	context.JSON(http.StatusCreated, content)
}
