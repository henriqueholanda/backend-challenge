package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OkResponse(context *gin.Context, content gin.H) {
	context.JSON(http.StatusOK, content)
}
